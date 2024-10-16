<script lang="ts">
	import { enhance } from '$app/forms';
	import {
		Button,
		Col,
		Container,
		Card,
		CardHeader,
		CardTitle,
		CardBody,
		Form,
		Input,
		Image,
		ListGroup,
		ListGroupItem,
		Row,
		Spinner,
		Toast
	} from '@sveltestrap/sveltestrap';
	import type { SubmitFunction } from '@sveltejs/kit';
	import { gql, request } from 'graphql-request';
	import { API_URL } from '$lib/constants';

	type WineClass = { [key: string]: number };

	let processing: boolean;
	let results: { name: string; probablity: number }[] = [];
	let imageURL = '';
	let selected = -1;
	let showToast = false;

	const authorizedExtensions = ['.jpg', '.jpeg', '.png'];

	const submitImage: SubmitFunction = () => {
		results = [];
		processing = true;
		return async ({ result }) => {
			processing = false;
			if (result.type === 'failure') {
				console.error('Failed to get results', result);
			}
			if (result.type === 'success') {
				if (!result.data?.results) {
					// TODO: Handle error responses
					console.error('No results available');
					return;
				}
				imageURL = result.data.results[0].name;
				const classes: WineClass = result.data.results[0].entities[0].classes;
				for (let c in classes) {
					results.push({ name: c, probablity: classes[c] });
				}
				results = results.sort((r) => r.probablity);
			}
		};
	};

	let formValues = {
		name: results[selected]?.name,
		grape: '',
		region: '',
		location: '',
		notes: '',
		image: imageURL,
		username: 'app_user'
	};

	const handleBookmarkSave = async () => {
		const document = gql`
			mutation {
				addBookmark(input: { 
				name: "${formValues.name}", 
				region: "${formValues.region}",
				grape: "${formValues.grape}",  
				location: "${formValues.location}",  
				notes:  "${formValues.notes}",
				image:  "${imageURL}",
				username:  "${formValues.username}",      
				}) {
					name
					region
					grape
					location
					notes
					image
					username
				}
			}`;
		const response: Response = await request(API_URL, document);
		if (!response.ok) {
			console.error(response.statusText);
			// TODO: show error message
			return;
		}
		// This also needs work
		showToast = true;
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
				<Button color="primary" type="submit">Submit</Button>
			</form>
		</CardBody>
	</Card>
</Container>
{#if results.length}
	<Container class="mt-5">
		<Row class="align-items-center">
			<h2 class="text-center">Results</h2>
			<Col sm="8">
				{#each results as result, i}
					<ListGroup>
						<ListGroupItem
							active={selected == i}
							on:click={() => {
								selected = i;
								formValues.name = results[selected].name;
							}}
							action>{result.name}</ListGroupItem
						>
					</ListGroup>
				{/each}
			</Col>
			<Col sm="4">
				<Image src={imageURL} alt="Uploaded Image" class="object-fit-contain mw-100" />
			</Col>
		</Row>
	</Container>
	{#if selected > -1}
		<Container class="mt-5">
			<Row class="align-items-center">
				<h2 class="text-center">Save Bookmark</h2>
				<Card>
					<Form class="p-5">
						<Input type="text" name="name" placeholder="Name" bind:value={formValues.name} />
						<Input type="text" name="grape" placeholder="Grape" bind:value={formValues.grape} />
						<Input type="text" name="region" placeholder="Region" bind:value={formValues.region} />
						<Input
							type="text"
							name="location"
							placeholder="Location"
							bind:value={formValues.location}
						/>
						<Input type="textarea" name="notes" placeholder="Notes" bind:value={formValues.notes} />
					</Form>
					<Button color="primary" type="submit" on:click={handleBookmarkSave}>Save</Button>
				</Card>
			</Row>
		</Container>
	{/if}
{:else if processing}
	<Container class="mt-5 d-flex justify-content-center">
		<Spinner />
	</Container>
{/if}

{#if showToast}
	<Toast isOpen={showToast} color="success" class="position-fixed bottom-0 end-0 m-3">
		<p class="mb-0">Bookmark saved successfully</p>
	</Toast>
{/if}
