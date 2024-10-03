import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';
import { generateV4UploadSignedUrl } from '$lib/gcs/generate-signed-url';
import { GCP_BASE_URL } from '$lib/constants';
import path from 'path';

export const actions: Actions = {
	uploadImage: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('image_upload') as File;
		// Want to prevent dupe file name uploads. Would probably do something smarter
		// like hash the image and use that as the file name. For now, just append a timestamp
		let fileName = path.parse(file.name).name;
		const fileExt = path.parse(file.name).ext;
		fileName = fileName + '-' + Date.now().toString() + fileExt;

		try {
			const url = await generateV4UploadSignedUrl(fileName);
			if (!url) {
				return fail(500, { error: 'No signed URL returned from generateV4UploadSignedUrl' });
			}

			const responseHeaders = new Headers();
			responseHeaders.append('Content-Type', 'application/octet-stream');

			const response = await fetch(url, {
				method: 'PUT',
				headers: responseHeaders,
				body: file
			});
			if (!response.ok) {
				return fail(response.status, { error: response.statusText });
			}
			return GCP_BASE_URL + fileName;
		} catch (error) {
			console.error('error', error);
		}
	}
};
