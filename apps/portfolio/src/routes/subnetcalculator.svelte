<script>
	import { onMount } from 'svelte';
	let cidr = '192.168.0.0/24';
	let subnets = [];

	onMount(() => {
		calc();
	});

	const calc = () => {
		const { octets, prefix } = parseCidr(cidr);
		const subnetInfo = getSubnetInfo(octets, prefix);
		subnets = [subnetInfo];
	};

	const parseCidr = (cidr) => {
		const [address, prefixStr] = cidr.split('/');
		const octets = address.split('.').map(Number);
		for (const octet of octets) {
			if (octet < 0) {
				return null;
			}
			if (octet > 255) {
				return null;
			}
		}
		const prefix = Number(prefixStr);
		if (prefix < 1) {
			return null;
		}
		if (prefix > 32) {
			return null;
		}
		return {
			octets,
			prefix
		};
	};

	const getSubnetInfo = (octets, prefix) => {
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
		return {
			networkCidr,
			networkMask,
			hosts,
			firstUsableAddress,
			lastUsableAddress,
			prefix,
			networkAddressOctets,
			broadcastAddressOctets
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
		const lowerSubnetInfo = getSubnetInfo(subnet.networkAddressOctets, subnet.prefix + 1);
		const upperSubnetInfo = getSubnetInfo(
			getNextAddressOctets(lowerSubnetInfo.broadcastAddressOctets),
			subnet.prefix + 1
		);
		subnets[index] = lowerSubnetInfo;
		subnets = subnets
			.slice(0, index + 1)
			.concat([upperSubnetInfo])
			.concat(subnets.slice(index + 1));
	};

	const join = (index) => {
		const subnet = subnets[index];
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
							{#if i > 0}
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
</div>
