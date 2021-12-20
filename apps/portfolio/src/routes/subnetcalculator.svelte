<script>
	import { v4 as uuidv4 } from 'uuid';
	import { onMount } from 'svelte';
	let cidr = '192.168.0.0/24';
	let subnets = [];
	let errorMsg = null;

	onMount(() => {
		calc();
	});

	const calc = () => {
		const { octets, prefix } = parseCidr(cidr);
		if (octets === null) {
			errorMsg = `${cidr} does not have a valid IP address.`;
			return;
		}
		if (prefix === null) {
			errorMsg = `${cidr} does not have a valid prefix.`;
			return;
		}
		errorMsg = null;
		const subnetInfo = getSubnetInfo(octets, prefix, null);
		subnets = [subnetInfo];
	};

	const parseCidr = (cidr) => {
		const [address, prefixStr] = cidr.split('/');
		let octets = address.split('.').map(Number);
		let octetCount = 0;
		for (const octet of octets) {
			if (Number.isNaN(octet)) {
				octets = null;
				break;
			}
			if (octet < 0) {
				octets = null;
				break;
			}
			if (octet > 255) {
				octets = null;
				break;
			}
			octetCount++;
		}
		if (octetCount !== 4) {
			octets = null;
		}
		let strOctets = address.split('.');
		for (const strOctet of strOctets) {
			if (strOctet === '') {
				octets = null;
				break;
			}
		}
		let prefix = Number(prefixStr);
		if (Number.isNaN(prefix)) {
			prefix = null;
		} else if (prefix < 1) {
			prefix = null;
		} else if (prefix > 32) {
			prefix = null;
		}
		return {
			octets,
			prefix
		};
	};

	const getSubnetInfo = (octets, prefix, uuid) => {
		const networkMaskOctets = getMaskOctets(prefix, false);
		const hostMaskOctets = getMaskOctets(prefix, true);
		const networkAddressOctets = getAddressOctets(octets, networkMaskOctets, false);
		const broadcastAddressOctets = getAddressOctets(octets, hostMaskOctets, true);
		const networkCidr = `${networkAddressOctets.join('.')}/${prefix}`;
		const networkMask = networkMaskOctets.join('.');
		let hosts = 2 ** (32 - prefix) - 2;
		let firstUsableAddress = getNextAddressOctets(networkAddressOctets).join('.');
		let lastUsableAddress = getPrevAddressOctets(broadcastAddressOctets).join('.');
		if (prefix === 32) {
			hosts = 1;
			firstUsableAddress = octets.join('.');
			lastUsableAddress = octets.join('.');
		} else if (prefix === 31) {
			hosts = 2;
			firstUsableAddress = networkAddressOctets.join('.');
			lastUsableAddress = getNextAddressOctets(networkAddressOctets).join('.');
		}
		if (uuid === null) {
			uuid = uuidv4();
		}
		const joined = [];
		return {
			uuid,
			networkCidr,
			networkMask,
			hosts,
			firstUsableAddress,
			lastUsableAddress,
			prefix,
			networkAddressOctets,
			broadcastAddressOctets,
			joined
		};
	};

	const getNextAddressOctets = (octets) => {
		const addressOctets = octets.slice(0, 3);
		addressOctets.push(octets[3] + 1);
		return addressOctets;
	};

	const getPrevAddressOctets = (octets) => {
		const addressOctets = octets.slice(0, 3);
		addressOctets.push(octets[3] - 1);
		return addressOctets;
	};

	const getAddressOctets = (octets, maskOctets, broadcast) => {
		const addressOctets = [];
		for (let i = 0; i < octets.length; i++) {
			const octet = octets[i];
			if (broadcast) {
				const maskOctet = maskOctets[i];
				addressOctets.push(octet | maskOctet);
			} else {
				const maskOctet = maskOctets[i];
				addressOctets.push(octet & maskOctet);
			}
		}
		return addressOctets;
	};

	const getMaskOctets = (prefix, host) => {
		const octets = [];
		let octetBin = '';
		for (let i = 0; i < 32; i++) {
			let addOne = prefix > i;
			if (host) {
				addOne = prefix <= i;
			}
			if (addOne) {
				octetBin = octetBin + '1';
			} else {
				octetBin = octetBin + '0';
			}
			if (octetBin.length === 8) {
				octets.push(parseInt(octetBin, 2));
				octetBin = '';
			}
		}
		return octets;
	};

	const divide = (index) => {
		const subnet = subnets[index];
		const lowerSubnetInfo = getSubnetInfo(
			subnet.networkAddressOctets,
			subnet.prefix + 1,
			subnet.uuid
		);
		lowerSubnetInfo.joined = lowerSubnetInfo.joined.concat(subnet.joined);
		const upperSubnetInfo = getSubnetInfo(
			getNextAddressOctets(lowerSubnetInfo.broadcastAddressOctets),
			subnet.prefix + 1,
			null
		);
		lowerSubnetInfo.joined.push(upperSubnetInfo.uuid);
		upperSubnetInfo.joined.push(lowerSubnetInfo.uuid);
		subnets[index] = lowerSubnetInfo;
		subnets = subnets
			.slice(0, index + 1)
			.concat([upperSubnetInfo])
			.concat(subnets.slice(index + 1));
	};

	const join = (index) => {
		const subnetInfo = subnets[index];
		const pairUuid = subnetInfo.joined.pop();
		const { pairIndex, pairSubnetInfo } = getPairObject(pairUuid);
		let subnetUuid = pairSubnetInfo.joined.pop();
		while (subnetUuid !== subnetInfo.uuid) {
			subnetUuid = pairSubnetInfo.joined.pop();
		}
		let lowerSubnetInfo = subnetInfo;
		let upperSubnetInfo = pairSubnetInfo;
		if (pairIndex < index) {
			lowerSubnetInfo = pairSubnetInfo;
			upperSubnetInfo = subnetInfo;
		}
		const joinedSubnetInfo = getSubnetInfo(
			lowerSubnetInfo.networkAddressOctets,
			upperSubnetInfo.prefix - 1,
			lowerSubnetInfo.uuid
		);
		joinedSubnetInfo.joined = joinedSubnetInfo.joined.concat(lowerSubnetInfo.joined);
		subnets[Math.min(index, pairIndex)] = joinedSubnetInfo;
		subnets = subnets
			.slice(0, Math.min(index, pairIndex) + 1)
			.concat(subnets.slice(Math.max(index, pairIndex) + 1));
	};

	const getPairObject = (uuid) => {
		for (let i = 0; i < subnets.length; i++) {
			if (uuid === subnets[i].uuid) {
				return { pairIndex: i, pairSubnetInfo: subnets[i] };
			}
		}
		return { pairIndex: null, pairSubnetInfo: null };
	};
</script>

<div class="container mx-auto">
	<div class="flex justify-center py-4">
		<div class="form-control">
			<div class="flex space-x-2">
				<input type="text" bind:value={cidr} class="w-full input input-primary input-bordered" />
				<button on:click={calc} class="btn btn-primary">Calc</button>
			</div>
		</div>
	</div>
	{#if errorMsg}
		<h1 class="bg-base-300 text-secondary text-xl text-center p-4">{errorMsg}</h1>
	{:else}
		<div class="overflow-x-auto">
			<table class="table w-full">
				<thead>
					<tr>
						<th />
						<th>CIDR</th>
						<th>Mask</th>
						<th>Hosts</th>
						<th>First Usable Address</th>
						<th>Last Usable Address</th>
						<th />
						<th />
					</tr>
				</thead>
				<tbody>
					{#each subnets as subnet, i}
						<tr>
							<td>{i + 1}</td>
							<td>{subnet.networkCidr}</td>
							<td>{subnet.networkMask}</td>
							<td>{subnet.hosts}</td>
							<td>{subnet.firstUsableAddress}</td>
							<td>{subnet.lastUsableAddress}</td>
							<td>
								{#if subnet.prefix < 32}
									<button
										on:click={() => {
											divide(i);
										}}
										class="btn btn-secondary">Divide</button
									>
								{/if}
							</td>
							<td>
								{#if subnet.joined.length > 0}
									<button
										on:click={() => {
											join(i);
										}}
										class="btn btn-accent">Join</button
									>
								{/if}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	{/if}
</div>
