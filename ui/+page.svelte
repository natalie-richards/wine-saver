<script lang="ts">
	import {
		Container,
		Col,
		Row,
		Card,
		CardHeader,
		CardTitle,
		CardBody,
		Form
	} from '@sveltestrap/sveltestrap';

	const authorizedExtensions = ['.jpg', '.jpeg', '.png'];
</script>

<Container class="d-flex justify-content-center mt-5">
	<Card>
		<CardHeader>
			<CardTitle>Upload Here</CardTitle>
		</CardHeader>
		<CardBody>
			<Form>
				<input type="file" placeholder="File" accept={authorizedExtensions.join(',')} required />
			</Form>
		</CardBody>
	</Card>
</Container>
