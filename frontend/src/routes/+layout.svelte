<script lang="ts">
	import '../app.css';
	import { isAuthenticated, user, clearAuth } from '$lib/auth';
	import { goto } from '$app/navigation';

	let { children } = $props();
	
	function handleLogout() {
		clearAuth();
		goto('/login');
	}
</script>

<div class="min-h-screen bg-gray-50">
	{#if $isAuthenticated}
		<nav class="bg-white shadow">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div class="flex justify-between h-16">
					<div class="flex">
						<div class="flex-shrink-0 flex items-center" >
							<a href="/" class="text-lg font-bold ">Airbnb for Stuff 📦</a>
						</div>
						<div class="hidden sm:ml-6 sm:flex sm:space-x-8">
							<a href="/items" class="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
								Browse Items
							</a>
							<a href="/items/my" class="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
								My Items
							</a>
						</div>
					</div>
					<div class="hidden sm:ml-6 sm:flex sm:items-center gap-2">
						{#if $user}
							<span class="text-sm text-gray-500 mr-4">Hello, {$user.first_name}</span>
						{/if}
						<button 
							onclick={handleLogout}
							class="border px-3 py-1 text-sm text-gray-700 hover:bg-gray-50"
						>
							Logout
						</button>
						<button 
							onclick={handleLogout}
							class="border px-3 py-1 text-sm text-white bg-black hover:bg-white hover:text-black"
						>
							My Rentals
						</button>
					</div>
				</div>
			</div>
		</nav>
	{/if}
	
	<main>
		{@render children()}
	</main>
</div>
