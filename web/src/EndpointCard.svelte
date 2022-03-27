<script>
    export let endpoint;
    export let getAll;

    async function handleDelete() {
        await fetch('http://localhost:5005/delete', {
			method: 'POST',
			mode: 'cors',
			body: JSON.stringify({ id: endpoint.id })
		}).catch(_ => {
            console.log("delete failed"); // TODO
		}).then(response => {
			if (response.ok) {
				getAll();
			}
		})
    }
</script>

<main>
    <div id="endpoint-card">
        <div id="endpoint-title">
            {endpoint.httpVerb} /{endpoint.path}
        </div>
        <div id="endpoint-id">
            <b>ID</b>: {endpoint.id}
        </div>
        <button id="endpoint-delete" on:click={() => handleDelete()}>Delete</button>
    </div>
</main>

<style>
    main {
        display: inline-block;
        margin: 0 auto;
        padding: 0;
        height: 200px;
        width: 33%;
    }

    #endpoint-card {
        display: inline-block;
        margin: 0;
        padding: 20px 40px 20px 40px;

        border-color: lightgrey;
        border-width: 2px;
        border-style: solid
    }

    #endpoint-title {
        margin-bottom: 10px;
        color: #0099f6;
        font-size: 20px;
    }

    #endpoint-id {
        margin-bottom: 20px;
    }

    button {
		color: white;
		background-color: red;
		text-transform: uppercase;
		border-width: 0px;
		font-weight: 500;
		height: 40px;
		width: 150px;
	}
</style>