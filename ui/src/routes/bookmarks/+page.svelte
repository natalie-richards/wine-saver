<script lang="ts">
	import { gql, request } from 'graphql-request';
	import { API_URL } from '$lib/constants';
	import type { Bookmark } from '$lib/types';
	import { Col, Container, Card, Image, Row } from '@sveltestrap/sveltestrap';

	let wineBookmarks: Bookmark[] = [];

	const document = gql`
		{
			listBookmarks {
				name
				notes
				region
				location
				grape
				image
			}
		}
	`;

	const getBookmarks = async (): Promise<Bookmark[]> => {
		const responseData = await request(API_URL, document).then((data) => data);
		return responseData.listBookmarks;
	};
</script>

<Container>
	{#await getBookmarks() then data}
		<Row>
			{#if data?.length}
				{#each data as bookmark}
					<Col sm="4">
						<Card>
							{#if bookmark.image}
								<Image src={bookmark.image} alt={bookmark.name} class="object-fit-contain mw-100" />
							{/if}
						</Card>
					</Col>
				{/each}
			{:else}
				<p>No bookmarks found</p>
			{/if}
		</Row>
	{/await}
</Container>
