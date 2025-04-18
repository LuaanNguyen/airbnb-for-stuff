<script lang="ts">
  import { onMount } from 'svelte';

  interface Item {
    id: number;
    name: string;
    description: string;
    image: string;
    price: number;
    date_listed: string;
    quantity: number;
    available: boolean;
  }

  let items: Item[] = [];
  let loading = true;
  let error: string | null = null;

  onMount(async () => {
    try {
      const response = await fetch('http://localhost:8080/api/items');
      if (!response.ok) throw new Error('Failed to fetch items');
      const data = await response.json();
      items = data.slice(0, 20); // Limit to 20 items
    } catch (e) {
      error = e instanceof Error ? e.message : 'An error occurred';
    } finally {
      loading = false;
    }
  });

  // Format price to USD
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD'
    }).format(price);
  };
</script>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <h1 class="text-3xl font-bold text-gray-900 mb-8">Available Items</h1>
  
  {#if loading}
    <div class="text-center py-12">
      <div class="text-gray-600">Loading items...</div>
    </div>
  {:else if error}
    <div class="text-center py-12">
      <div class="text-red-600">{error}</div>
    </div>
  {:else}
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      {#each items as item}
        {#if item.available}
          <div class="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-200 hover:-translate-y-1 hover:shadow-lg">
            <div class="aspect-w-16 aspect-h-9 bg-gray-200">
              {#if item.image}
                <img
                  src={`data:image/jpeg;base64,${item.image}`}
                  alt={item.name}
                  class="w-full h-48 object-cover"
                />
              {:else}
                <div class="w-full h-48 flex items-center justify-center bg-gray-100 text-gray-500">
                  No image available
                </div>
              {/if}
            </div>
            
            <div class="p-4">
              <h2 class="text-xl font-semibold text-gray-900 mb-2 truncate">
                {item.name}
              </h2>
              
              <p class="text-2xl font-bold text-green-600 mb-2">
                {formatPrice(item.price)}
              </p>
              
              <p class="text-gray-600 text-sm mb-3 line-clamp-2">
                {item.description}
              </p>
              
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-500">
                  Available: {item.quantity}
                </span>
                <button class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md text-sm font-medium transition-colors duration-200">
                  View Details
                </button>
              </div>
            </div>
          </div>
        {/if}
      {/each}
    </div>
  {/if}
</div> 