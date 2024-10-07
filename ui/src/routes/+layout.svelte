<script lang="ts">
	import { Styles, Navbar, Nav, NavItem, NavLink } from '@sveltestrap/sveltestrap';
</script>

<Styles />
<Navbar color="light" light expand="md">
	<Nav class="ms-auto" navbar>
		<NavItem>
			<NavLink href="/">Upload</NavLink>
		</NavItem>
		<NavItem>
			<NavLink href="/bookmarks">My Bookmarks</NavLink>
		</NavItem>
	</Nav>
</Navbar>
<slot />
