import type { Actions } from './$types';
import { fail } from '@sveltejs/kit';
import { generateV4UploadSignedUrl } from '$lib/gcs/generate-signed-url';
import { GCP_BASE_URL, WINE_LABEL_API_URL } from '$lib/constants';
import { RAPIDAPI_KEY } from '$env/static/private';
import path from 'path';

export const actions: Actions = {
	uploadImageAndGetWineResult: async ({ request }) => {
		const formData = await request.formData();
		const file = formData.get('image_upload') as File;
		// Want to prevent dupe file uploads. Would probably do something smarter
		// like hash the image and use that as the file name. For now, just append a timestamp
		// Also I should be compressing the image to save on storage costs
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

			// TODO: Compress file before upload

			const response = await fetch(url, {
				method: 'PUT',
				headers: responseHeaders,
				body: file
			});
			if (!response.ok) {
				return fail(response.status, { error: response.statusText });
			}
		} catch (error) {
			console.error('error', error);
		}

		// Call the wine label API
		const data = new FormData();
		data.append('url', GCP_BASE_URL + fileName);
		const options = {
			method: 'POST',
			headers: {
				'X-RapidAPI-Key': RAPIDAPI_KEY
			},
			body: data
		};

		try {
			const response = await fetch(WINE_LABEL_API_URL, options);
			if (!response.ok) {
				return fail(response.status, { error: response.statusText });
			}
			const result = await response.json();
			return result;
		} catch (error) {
			console.error(error);
		}
	}
};
