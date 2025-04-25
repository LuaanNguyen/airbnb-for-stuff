<script lang="ts">
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { isAuthenticated } from '$lib/auth';

  onMount(() => {
    if ($isAuthenticated) {
      goto('/items');
    } else {
      goto('/login');
    }
  });
</script>

<div class="flex items-center justify-center h-screen">
  <div class="text-center">
    <h1 class="text-2xl font-bold mb-4">Welcome to Airbnb for Stuff</h1>
    <p class="text-gray-600 mb-6">Redirecting you to the right place...</p>
    <div class="animate-spin h-8 w-8 border-4 border-blue-500 rounded-full border-t-transparent mx-auto"></div>
  </div>
</div>

  

  