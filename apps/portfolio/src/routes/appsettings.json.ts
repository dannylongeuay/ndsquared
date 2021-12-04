export async function get(): Promise<{ body: any }> {
	return {
		body: {
			API_BASE_URL: process.env['API_BASE_URL']
		}
	};
}
