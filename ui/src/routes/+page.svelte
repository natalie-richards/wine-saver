<script lang="ts">
	import { enhance } from '$app/forms';
	import {
		Button,
		Container,
		Card,
		CardHeader,
		CardTitle,
		CardBody,
		Spinner
	} from '@sveltestrap/sveltestrap';

	import type { SubmitFunction } from '@sveltejs/kit';

	let processing: boolean;

	const authorizedExtensions = ['.jpg', '.jpeg', '.png'];

	const submitImage: SubmitFunction = () => {
		processing = true;
		return async ({ result }) => {
			processing = false;
			if (result.type === 'failure') {
				console.error('Failed to upload image', result);
			}
			if (result.type === 'success') {
				if (!result.data?.results) {
					// TODO: Handle error responses
					console.error('No results available');
					return;
				}
				const entities = result.data.results[0].entities;
				for (const entity of entities) {
					console.log(entity.classes);
				}
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
				action="?/uploadImageAndGetWineResult"
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
				{#if processing}
					<Spinner />
				{/if}
				<Button color="primary" type="submit">Submit</Button>
			</form>
		</CardBody>
	</Card>
</Container>
