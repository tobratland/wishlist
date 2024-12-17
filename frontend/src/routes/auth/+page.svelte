<script>
    import { goto } from "$app/navigation";

    let email = "";
    let password = "";
    let name = "";

    async function register() {
        const res = await fetch("/api/register", {
            method: "POST",
            body: JSON.stringify({ name, email, password }),
            headers: { "Content-Type": "application/json" },
        });
        if (res.ok) goto("/wishlists");
    }

    async function login() {
        const res = await fetch("/api/login", {
            method: "POST",
            body: JSON.stringify({ email, password }),
            headers: { "Content-Type": "application/json" },
        });
        if (res.ok) goto("/wishlists");
    }
</script>

<div class="min-h-screen bg-christmasGreen p-8">
    <div class="text-white text-3xl mb-4">
        🎄 Welcome to the Christmas Wishlist 🎅
    </div>

    <!-- Login Form -->
    <form class="mb-8 space-y-4" on:submit|preventDefault={login}>
        <input placeholder="Email" bind:value={email} class="block p-2" />
        <input
            type="password"
            placeholder="Password"
            bind:value={password}
            class="block p-2"
        />
        <button type="submit" class="bg-christmasRed p-2 text-white"
            >Login</button
        >
    </form>

    <!-- Register Form -->
    <form class="space-y-4" on:submit|preventDefault={register}>
        <input placeholder="Name" bind:value={name} class="block p-2" />
        <input placeholder="Email" bind:value={email} class="block p-2" />
        <input
            type="password"
            placeholder="Password"
            bind:value={password}
            class="block p-2"
        />
        <button type="submit" class="bg-christmasRed p-2 text-white"
            >Register</button
        >
    </form>
</div>
