<script>
    import { onMount } from "svelte";

    /**
     * @type {any[]}
     */
    let wishlists = [];
    let title = "";
    let description = "";

    async function fetchWishlists() {
        const res = await fetch("/api/wishlists", {
            headers: { Authorization: "Bearer YOUR_TOKEN" },
        });
        wishlists = await res.json();
    }

    async function createWishlist() {
        await fetch("/api/wishlists", {
            method: "POST",
            body: JSON.stringify({ title, description }),
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer YOUR_TOKEN",
            },
        });
        fetchWishlists();
    }

    onMount(fetchWishlists);
</script>

<div class="min-h-screen bg-snowfall p-8 text-gray-800">
    <h1 class="text-3xl text-christmasRed">🎁 My Wishlists 🎄</h1>

    <form class="space-y-2 my-4" on:submit|preventDefault={createWishlist}>
        <input placeholder="Title" bind:value={title} class="block p-2" />
        <textarea
            placeholder="Description"
            bind:value={description}
            class="block p-2"
        ></textarea>
        <button class="bg-christmasGreen text-white p-2">Add Wishlist</button>
    </form>

    <ul>
        {#each wishlists as wishlist}
            <li class="bg-white p-2 my-2 rounded shadow">
                <h2 class="text-christmasGreen">{wishlist.title}</h2>
                <p>{wishlist.description}</p>
            </li>
        {/each}
    </ul>
</div>
