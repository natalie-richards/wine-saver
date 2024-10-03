import type { Actions } from './$types';
import { generateV4UploadSignedUrl } from '$lib/gcs/generate-signed-url';

export const actions: Actions = {
	uploadImage: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('image_upload') as File;

		try {
			const url = await generateV4UploadSignedUrl(file.name);
			if (!url) {
				throw new Error('No signed URL returned from generateV4UploadSignedUrl');
			}

			const responseHeaders = new Headers();
			responseHeaders.append('Content-Type', 'application/octet-stream');

			const response = await fetch(url, {
				method: 'PUT',
				headers: responseHeaders
			});
			if (!response.ok) {
				throw new Error(response.statusText);
			}
		} catch (error) {
			console.error('error', error);
		}
	}
};
