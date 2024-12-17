<script>
    import { goto } from '$app/navigation';
    import { register } from '$lib/api';
  
    let name = '';
    let email = '';
    let password = '';
    let error = '';
  
    async function handleRegister() {
      error = '';
      const result = await register(name, email, password);
      if (result.success) {
        goto('/auth/login');
      } else {
        error = result.message || 'Registration failed. Try again.';
      }
    }
  </script>
  
  <div class="min-h-screen bg-snowfall flex flex-col items-center justify-center">
    <h1 class="text-3xl text-christmasRed mb-4">✨ Register 🎄</h1>
  
    {#if error}
      <p class="text-red-500">{error}</p>
    {/if}
  
    <form class="space-y-4" on:submit|preventDefault={handleRegister}>
      <input placeholder="Name" bind:value={name} class="block p-2 border rounded" required />
      <input placeholder="Email" bind:value={email} class="block p-2 border rounded" required />
      <input
        type="password"
        placeholder="Password"
        bind:value={password}
        class="block p-2 border rounded"
        required
      />
      <button class="bg-christmasGreen text-white px-4 py-2">Register</button>
    </form>
  </div>
  