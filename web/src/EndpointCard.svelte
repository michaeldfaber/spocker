<script>
    export let endpoint;
    export let getAll;

    async function handleDelete() {
        await fetch('http://0.0.0.0:5001/delete', {
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

<main id="endpoint-card">
    <div id="endpoint-verb">
        {endpoint.httpVerb}
    </div>
    <div id="endpoint-path">
        /{endpoint.path}
    </div>
    <button id="endpoint-delete" on:click={() => handleDelete()}>Delete</button>
</main>

<style>
    #endpoint-card {
        display: inline-block;
        margin: 0px 10px 20px 10px;
        padding: 10px 20px;
        min-width: 25%;

        border-color: lightgrey;
        border-width: 1px;
        border-style: solid;
    }

    #endpoint-verb {
        margin-bottom: 10px;
        font-size: 16px;
    }

    #endpoint-path {
        margin-bottom: 10px;
        font-size: 16px;
        font-family: monospace;
        color: #0099f6;
    }

    button {
		color: white;
		background-color: red;
		text-transform: uppercase;
		border-width: 0px;
		font-weight: 400;
		height: 40px;
		width: 150px;
	}
</style>