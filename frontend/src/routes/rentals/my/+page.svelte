<script lang="ts">
  import { onMount } from 'svelte';
  import { getMyRentals } from '$lib/services/api';
  import type { RentalWithDetails } from '$lib/types';
  import { isAuthenticated } from '$lib/auth';
  import { goto } from '$app/navigation';
  
  let rentals: RentalWithDetails[] = [];
  let loading = true;
  let error = '';
  
  onMount(async () => {
    // Check if user is authenticated
    if (!$isAuthenticated) {
      goto('/login');
      return;
    }
    
    try {
      rentals = await getMyRentals();
      loading = false;
    } catch (err) {
      error = err instanceof Error ? err.message : 'Failed to load your rentals';
      loading = false;
    }
  });
  
  // Format date for display
  function formatDate(dateString: string): string {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  }
  
  // Get status badge color
  function getStatusColor(status: string): string {
    switch (status.toLowerCase()) {
      case 'approved': return 'bg-green-100 text-green-800';
      case 'pending': return 'bg-yellow-100 text-yellow-800';
      case 'rejected': return 'bg-red-100 text-red-800';
      case 'completed': return 'bg-blue-100 text-blue-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  }
</script>

<div class="container mx-auto px-4 py-8">
  <h1 class="text-2xl font-bold mb-6">My Rental Requests</h1>
  
  {#if loading}
    <div class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-gray-900"></div>
    </div>
  {:else if error}
    <div class="bg-red-50 border border-red-200 text-red-800 p-4 rounded-md">
      <p>{error}</p>
    </div>
  {:else if rentals.length === 0}
    <div class="bg-gray-50 border border-gray-200 p-6 rounded-md text-center">
      <p class="text-lg text-gray-600">You don't have any rental requests yet.</p>
      <a href="/items" class="inline-block mt-4 px-4 py-2 bg-black text-white hover:bg-gray-800">
        Browse Items to Rent
      </a>
    </div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {#each rentals as rental}
        <div class="border rounded-lg overflow-hidden hover:shadow-md transition-shadow">
          <div class="p-4 border-b">
            <h2 class="text-xl font-semibold">{rental.item_name}</h2>
            <p class="text-gray-600 mt-1">{rental.description.substring(0, 100)}...</p>
          </div>
          
          <div class="p-4">
            <div class="flex justify-between items-center mb-2">
              <span class="text-gray-600">From Owner:</span>
              <span class="font-medium">{rental.owner_name}</span>
            </div>
            
            <div class="flex justify-between items-center mb-2">
              <span class="text-gray-600">Start Date:</span>
              <span>{formatDate(rental.start_date)}</span>
            </div>
            
            <div class="flex justify-between items-center mb-2">
              <span class="text-gray-600">End Date:</span>
              <span>{formatDate(rental.end_date)}</span>
            </div>
            
            <div class="flex justify-between items-center mb-2">
              <span class="text-gray-600">Price:</span>
              <span class="font-semibold">${(rental.total_price / 100).toFixed(2)}</span>
            </div>
            
            <div class="flex justify-between items-center mb-2">
              <span class="text-gray-600">Status:</span>
              <span class={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(rental.status)}`}>
                {rental.status.toUpperCase()}
              </span>
            </div>
          </div>
          
          <div class="p-4 bg-gray-50">
            <a href={`/items/${rental.item_id}`} class="text-black hover:underline block text-center">
              View Item Details
            </a>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div> 