<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { getItem, createRentalRequest } from '$lib/services/api';
  import { isAuthenticated } from '$lib/auth';
  import type { Item, RentalRequest } from '$lib/types';

  let item: Item | null = null;
  let loading = true;
  let error: string | null = null;
  let successMessage: string | null = null;

  // Rental form data
  let startDate = '';
  let endDate = '';
  let calculatedPrice = 0;
  let formSubmitting = false;

  // Get item ID from route params
  const itemId = parseInt($page.params.id);

  onMount(async () => {
    // Check if user is authenticated
    if (!$isAuthenticated) {
      goto('/login');
      return;
    }

    await loadItemDetails();
  });

  async function loadItemDetails() {
    try {
      loading = true;
      item = await getItem(itemId);
      loading = false;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to load item details';
      loading = false;
    }
  }

  // Calculate rental price based on days
  function calculatePrice() {
    if (!item || !startDate || !endDate) {
      calculatedPrice = 0;
      return;
    }

    const start = new Date(startDate);
    const end = new Date(endDate);

    // Return 0 if dates are invalid
    if (isNaN(start.getTime()) || isNaN(end.getTime()) || end <= start) {
      calculatedPrice = 0;
      return;
    }

    // Calculate days (including partial days)
    const days = Math.ceil((end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24));
    
    // Calculate price (item price is daily rate in cents)
    calculatedPrice = days * item.price;
  }

  // Watch for changes in dates to recalculate price
  $: if (startDate && endDate) {
    calculatePrice();
  }

  // Handle form submission
  async function handleRentalSubmit() {
    if (!item || !startDate || !endDate || calculatedPrice <= 0) {
      error = 'Please select valid rental dates';
      return;
    }

    try {
      formSubmitting = true;
      error = null;
      
      const rentalData: RentalRequest = {
        item_id: itemId,
        start_date: startDate,
        end_date: endDate,
        total_price: calculatedPrice
      };

      await createRentalRequest(rentalData);
      successMessage = 'Rental request submitted successfully!';
      
      // Reset form after successful submission
      startDate = '';
      endDate = '';
      calculatedPrice = 0;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to submit rental request';
    } finally {
      formSubmitting = false;
    }
  }

  // Format price to USD
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2
    }).format(price / 100); // Converting cents to dollars
  };
</script>

<div class="max-w-4xl mx-auto px-4 py-8">
  <div class="mb-4">
    <button 
      on:click={() => goto('/items')}
      class=" flex items-center cursor-pointer "
    >
      ← Back to Items
    </button>
  </div>
  
  {#if loading}
    <div class="text-center py-12">
      <div class="text-gray-600">Loading item details...</div>
    </div>
  {:else if error}
    <div class="bg-red-50 border border-red-200 text-red-800 p-4  mb-6">
      {error}
    </div>
  {:else if item}
    <div class="bg-white overflow-hidden">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 p-6">
        <!-- Item Image and Details -->
        <div>
          <div class="bg-gray-100 h-64 flex items-center justify-center text-gray-500  mb-4">
            No image available
          </div>
          
          <h1 class="text-3xl font-bold text-gray-900 mb-2">{item.name}</h1>
          <p class="text-xl font-bold text-green-600 mb-4">{formatPrice(item.price)} per day</p>
          
          <div class="prose prose-sm max-w-none mb-6">
            <p>{item.description}</p>
          </div>
          
          <div class="flex flex-wrap gap-4 text-sm text-gray-600">
            <div>
              <span class="font-medium">Category:</span> 
              <span>{item.category_id}</span>
            </div>
            <div>
              <span class="font-medium">Quantity:</span> 
              <span>{item.quantity} available</span>
            </div>
          </div>
        </div>
        
        <!-- Rental Form -->
        <div class="bg-gray-50 p-6 ">
          <h2 class="text-xl font-bold text-gray-900 mb-4">Rent This Item</h2>
          
          {#if successMessage}
            <div class="bg-green-50 border border-green-200 text-green-800 p-4 rounded-md mb-4">
              {successMessage}
            </div>
          {/if}
          
          <form on:submit|preventDefault={handleRentalSubmit} class="space-y-4">
            <div>
              <label for="start-date" class="block text-sm font-medium text-gray-700 mb-1">
                Start Date
              </label>
              <input
                type="date"
                id="start-date"
                bind:value={startDate}
                min={new Date().toISOString().split('T')[0]}
                required
                class="w-full px-3 py-2 border border-gray-300 "
              />
            </div>
            
            <div>
              <label for="end-date" class="block text-sm font-medium text-gray-700 mb-1">
                End Date
              </label>
              <input
                type="date"
                id="end-date"
                bind:value={endDate}
                min={startDate || new Date().toISOString().split('T')[0]}
                required
                class="w-full px-3 py-2 border border-gray-300 "
              />
            </div>
            
            {#if calculatedPrice > 0}
              <div class="bg-blue-50 p-4 ">
                <h3 class="font-semibold text-blue-800 mb-2">Rental Summary</h3>
                <div class="flex justify-between">
                  <span>Total Price:</span>
                  <span class="font-bold">{formatPrice(calculatedPrice)}</span>
                </div>
              </div>
            {/if}
            
            <button
              type="submit"
              disabled={formSubmitting || !startDate || !endDate || calculatedPrice <= 0}
              class="w-full py-3 bg-blue-600 text-white  font-medium hover:bg-blue-700 disabled:bg-gray-400"
            >
              {formSubmitting ? 'Submitting...' : 'Submit Rental Request'}
            </button>
          </form>
        </div>
      </div>
    </div>
  {:else}
    <div class="text-center py-12">
      <div class="text-gray-600">Item not found</div>
    </div>
  {/if}
</div> 