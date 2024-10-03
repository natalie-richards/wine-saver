import { BUCKET_NAME } from '../constants';
import { Storage, type GetSignedUrlConfig } from '@google-cloud/storage';

// Creates a client
const storage = new Storage({
	projectId: 'vast-watch-437017-f9',
	keyFilename: 'src/lib/gcs/service-account-key.json'
});

export async function generateV4UploadSignedUrl(fileName: string) {
	const options: GetSignedUrlConfig = {
		version: 'v4',
		action: 'write',
		expires: Date.now() + 15 * 60 * 1000, // 15 minutes,
		contentType: 'application/octet-stream'
	};

	// Get a v4 signed URL for uploading file
	const [url] = await storage.bucket(BUCKET_NAME).file(fileName).getSignedUrl(options);

	return url;
}
