import { BUCKET_NAME } from '../constants';
// Imports the Google Cloud client library
import { Storage, type GetSignedUrlConfig } from '@google-cloud/storage';

// Creates a client
const storage = new Storage({
	projectId: 'vast-watch-437017-f9',
	keyFilename: 'src/lib/gcs/service-account-key.json'
});

export async function generateV4UploadSignedUrl(fileName: string) {
	// These options will allow temporary uploading of the file with outgoing
	// Content-Type: application/octet-stream header.
	const options: GetSignedUrlConfig = {
		version: 'v4',
		action: 'write',
		expires: Date.now() + 15 * 60 * 1000, // 15 minutes
		contentType: 'application/octet-stream'
	};

	// Get a v4 signed URL for uploading file
	const [url] = await storage.bucket(BUCKET_NAME).file(fileName).getSignedUrl(options);

	console.log('Generated PUT signed URL:');
	console.log(url);

	return url;
}
