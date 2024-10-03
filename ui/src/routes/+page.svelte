<script lang="ts">
	import { enhance } from '$app/forms';
	import {
		Button,
		Container,
		Card,
		CardHeader,
		CardTitle,
		CardBody
	} from '@sveltestrap/sveltestrap';
	// import { gql, request } from 'graphql-request';
	// import { API_URL } from '$lib/constants';
	// import { onMount } from 'svelte';
	import type { SubmitFunction } from '@sveltejs/kit';
	// let processing: boolean;

	// const document = gql`
	// 	{
	// 		listBookmarks {
	// 			name
	// 			region
	// 		}
	// 	}
	// `;

	// const test = async () => await request(API_URL, document);

	// onMount(() => {
	// 	test().then((data) => console.log(data));
	// });

	const authorizedExtensions = ['.jpg', '.jpeg', '.png'];

	const submitImage: SubmitFunction = () => {
		return async ({ result }) => {
			if (result.type === 'failure') {
				console.error('Failed to upload image', result);
			}
			if (result.type === 'success') {
				console.log('Image uploaded pub url', result.data);
			}
		};
	};
</script>

<Container class="d-flex justify-content-center mt-5">
	<Card>
		<CardHeader>
			<CardTitle>Upload Here</CardTitle>
		</CardHeader>
		<CardBody>
			<form
				action="?/uploadImage"
				method="POST"
				enctype="multipart/form-data"
				use:enhance={submitImage}
			>
				<input
					type="file"
					name="image_upload"
					placeholder="File"
					accept={authorizedExtensions.join(',')}
					required
				/>
				<Button color="primary" type="submit">Submit</Button>
			</form>
		</CardBody>
	</Card>
</Container>
