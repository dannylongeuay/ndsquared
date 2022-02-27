export async function get(): Promise<{ body: any }> {
	return {
		body: {
			API_BASE_URL: process.env['API_BASE_URL'],
			GOAPI_BASE_URL: process.env['GOAPI_BASE_URL'],
			AUTH0_DOMAIN: process.env['AUTH0_DOMAIN'],
			AUTH0_CLIENT_ID: process.env['AUTH0_CLIENT_ID'],
			AUTH0_CLIENT_SECRET: process.env['AUTH0_CLIENT_SECRET'],
			AUTH0_API_AUDIENCE: process.env['AUTH0_API_AUDIENCE']
		}
	};
}
