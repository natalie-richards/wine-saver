<script lang="ts">
	import { gql, request } from 'graphql-request';
	import { API_URL } from '$lib/constants';
	import type { Bookmark } from '$lib/types';
	import { Col, Container, Card, Image, Row } from '@sveltestrap/sveltestrap';

	type wineResponse = { listBookmarks: Bookmark[] };
	let bookmarks: Bookmark[] = [];

	const document = gql`
		{
			listBookmarks {
				name
				region
				notes
				location
				grape
				image
				username
			}
		}
	`;

	const getBookmarks = async (): Promise<Bookmark[]> => {
		await request(API_URL, document)
			.then((data) => (bookmarks = (data as wineResponse)?.listBookmarks))
			.catch((err) => console.error(err));

		return bookmarks;
	};
</script>

<Container class="mt-5">
	{#await getBookmarks() then data}
		<Row>
			{#if data?.length}
				{#each data as bookmark}
					<Col sm="6">
						<Card class="m-2">
							<Row>
								{#if bookmark.image}
									<Col sm="6" class="text-center" style="max-height: 30vh">
										<Image
											src={bookmark.image}
											alt={bookmark.name}
											class="object-fit-cover mw-100 mh-100"
										/>
									</Col>
								{/if}
								<Col class="d-flex align-items-center flex-wrap">
									<h5>{bookmark.name}</h5>
									<span><b>Grape:</b> {bookmark.grape}</span>
									<span><b>Region:</b> {bookmark.region}</span>
									<span><b>Purchased:</b> {bookmark.location}</span>
									<span><b>Notes:</b> {bookmark.notes}</span>
								</Col>
							</Row>
						</Card>
					</Col>
				{/each}
			{:else}
				<p>No bookmarks found</p>
			{/if}
		</Row>
	{/await}
</Container>
