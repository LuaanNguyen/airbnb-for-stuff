<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { getMyItems, createItem, deleteItem, updateItem, getAllCategories } from '$lib/services/api';
  import { isAuthenticated, user } from '$lib/auth';
  import type { Item, Category } from '$lib/types';

  let myItems: Item[] = [];
  let categories: Category[] = [];
  let loading = true;
  let error: string | null = null;
  let successMessage: string | null = null;

  // Form data for new/edit item
  let formMode: 'add' | 'edit' = 'add';
  let currentItemId: number | null = null;
  let showForm = false;
  let formSubmitting = false;
  let itemName = '';
  let itemDescription = '';
  let itemCategoryId: number | null = null;
  let itemPrice = 0;
  let itemQuantity = 1;
  let itemAvailable = true;

  onMount(async () => {
    // Check if user is authenticated
    if (!$isAuthenticated) {
      goto('/login');
      return;
    }

    try {
      loading = true;
      // Load user's items and categories
      const [userItems, cats] = await Promise.all([
        getMyItems(),
        getAllCategories()
      ]);

      categories = cats;
      
      myItems = userItems;
      
      loading = false;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to load data';
      loading = false;
    }
  });

  // Reset form fields
  function resetForm() {
    formMode = 'add';
    currentItemId = null;
    itemName = '';
    itemDescription = '';
    itemCategoryId = null;
    itemPrice = 0;
    itemQuantity = 1;
    itemAvailable = true;
  }

  // Show form for adding a new item
  function showAddForm() {
    resetForm();
    showForm = true;
  }

  // Show form for editing an existing item
  function showEditForm(item: Item) {
    formMode = 'edit';
    currentItemId = item.id || null;
    itemName = item.name;
    itemDescription = item.description;
    itemCategoryId = item.category_id;
    itemPrice = item.price;
    itemQuantity = item.quantity;
    itemAvailable = item.available;
    showForm = true;
  }

  // Handle form submission for add/edit
  async function handleSubmit() {
    if (!itemName || !itemDescription || !itemCategoryId) {
      error = 'Please fill out all required fields';
      return;
    }

    try {
      formSubmitting = true;
      error = null;
      successMessage = null;
      
      const itemData: Item = {
        name: itemName,
        description: itemDescription,
        category_id: itemCategoryId,
        price: itemPrice,
        quantity: itemQuantity,
        available: itemAvailable
      };

      if (formMode === 'add') {
        // Create new item
        const newItem = await createItem(itemData);
        myItems = [...myItems, newItem];
        successMessage = 'Item created successfully!';
      } else if (formMode === 'edit' && currentItemId) {
        // Update existing item
        const updatedItem = await updateItem(currentItemId, itemData);
        myItems = myItems.map(item => 
          item.id === currentItemId ? updatedItem : item
        );
        successMessage = 'Item updated successfully!';
      }

      // Hide form after successful submission
      showForm = false;
      resetForm();
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to save item';
    } finally {
      formSubmitting = false;
    }
  }

  // Handle item deletion
  async function handleDelete(id: number) {
    if (!confirm('Are you sure you want to delete this item?')) {
      return;
    }

    try {
      error = null;
      successMessage = null;
      
      await deleteItem(id);
      myItems = myItems.filter(item => item.id !== id);
      successMessage = 'Item deleted successfully!';
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to delete item';
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

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="flex justify-between items-center mb-8">
    <h1 class="text-3xl font-bold text-gray-900">My Items</h1>
    <button
      on:click={showAddForm}
      class="px-4 py-2 bg-black text-white hover:bg-white hover:text-black border"
    >
      Add New Item
    </button>
  </div>
  
  {#if error}
    <div class="bg-red-50 border border-red-200 text-red-800 p-4 rounded-md mb-6">
      {error}
    </div>
  {/if}
  
  {#if successMessage}
    <div class="bg-green-50 border border-green-200 text-green-800 p-4 rounded-md mb-6">
      {successMessage}
    </div>
  {/if}
  
  {#if showForm}
    <div class="bg-white rounded-lg shadow-md p-6 mb-8">
      <h2 class="text-xl font-bold mb-4">{formMode === 'add' ? 'Add New Item' : 'Edit Item'}</h2>
      
      <form on:submit|preventDefault={handleSubmit} class="space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <!-- Item Name -->
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700 mb-1">
              Item Name *
            </label>
            <input
              type="text"
              id="name"
              bind:value={itemName}
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md"
            />
          </div>
          
          <!-- Category -->
          <div>
            <label for="category" class="block text-sm font-medium text-gray-700 mb-1">
              Category *
            </label>
            <select
              id="category"
              bind:value={itemCategoryId}
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md"
            >
              <option value={null}>Select a category</option>
              {#each categories as category}
                <option value={category.id}>{category.name}</option>
              {/each}
            </select>
          </div>
          
          <!-- Price -->
          <div>
            <label for="price" class="block text-sm font-medium text-gray-700 mb-1">
              Price (cents) *
            </label>
            <input
              type="number"
              id="price"
              bind:value={itemPrice}
              min="0"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md"
            />
            <p class="text-sm text-gray-500 mt-1">
              {formatPrice(itemPrice)} per day
            </p>
          </div>
          
          <!-- Quantity -->
          <div>
            <label for="quantity" class="block text-sm font-medium text-gray-700 mb-1">
              Quantity *
            </label>
            <input
              type="number"
              id="quantity"
              bind:value={itemQuantity}
              min="1"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md"
            />
          </div>
        </div>
        
        <!-- Description -->
        <div>
          <label for="description" class="block text-sm font-medium text-gray-700 mb-1">
            Description *
          </label>
          <textarea
            id="description"
            bind:value={itemDescription}
            rows="4"
            required
            class="w-full px-3 py-2 border border-gray-300 rounded-md"
          ></textarea>
        </div>
        
        <!-- Available -->
        <div class="flex items-center">
          <input
            type="checkbox"
            id="available"
            bind:checked={itemAvailable}
            class="h-4 w-4 text-blue-600 border-gray-300 rounded"
          />
          <label for="available" class="ml-2 block text-sm text-gray-700">
            Item is available for rent
          </label>
        </div>
        
        <div class="flex space-x-3 pt-4">
          <button
            type="button"
            on:click={() => showForm = false}
            class="px-4 py-2 border border-gray-300 text-gray-700 hover:bg-gray-50"
          >
            Cancel
          </button>
          <button
            type="submit"
            disabled={formSubmitting}
            class="px-4 py-2  bg-black text-white hover:bg-white hover:text-black border"
          >
            {formSubmitting ? 'Saving...' : 'Save Item'}
          </button>
        </div>
      </form>
    </div>
  {/if}
  
  {#if loading}
    <div class="text-center py-12">
      <div class="text-gray-600">Loading your items...</div>
    </div>
  {:else if myItems.length === 0}
    <div class="bg-gray-50 p-8 text-center ">
      <p class="text-gray-600 mb-4">You don't have any items listed yet.</p>
      <button
        on:click={showAddForm}
        class="px-4 py-2 bg-black text-white hover:bg-white hover:text-black border"
      >
        Add Your First Item
      </button>
    </div>
  {:else}
    <div class="bg-white shadow-md  overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Item
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Price
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Quantity
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Status
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          {#each myItems as item}
            <tr>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-gray-900">{item.name}</div>
                <div class="text-sm text-gray-500 truncate max-w-xs">{item.description}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{formatPrice(item.price)}</div>
                <div class="text-xs text-gray-500">per day</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {item.quantity}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                {#if item.available}
                  <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                    Available
                  </span>
                {:else}
                  <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-gray-100 text-gray-800">
                    Unavailable
                  </span>
                {/if}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  on:click={() => showEditForm(item)}
                  class="text-blue-600 hover:text-blue-900 mr-3"
                >
                  Edit
                </button>
                <button
                  on:click={() => item.id && handleDelete(item.id)}
                  class="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</div> 