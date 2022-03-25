<script>
	let httpVerb = "GET"
	let endpoint;
	let expectedJsonResponse;
	let createMessage = "";

	function handleHttpVerb(e) {
		httpVerb = e.target.value;
	}

	function handleEndpoint(e) {
		endpoint = e.target.value;
	}

	function handleExpectedJsonResponse(e) {
		expectedJsonResponse = e.target.value;
	}

	async function handleCreate() {
		const expectedJsonResponseParsed = JSON.parse(expectedJsonResponse)
		const res = await fetch('http://localhost:5005/create', {
			method: 'POST',
			mode: 'cors',
			body: JSON.stringify({
				httpVerb: httpVerb,
				endpoint: endpoint,
				expectedJsonResponse: expectedJsonResponseParsed
			})
		}).then(response => {
			console.log(response);
			if (response.ok) {
				createMessage = "Successfully created " + httpVerb + " " + endpoint + " endpoint!";
			} else {
				createMessage = "Create failed."
			}
		})
	}
</script>

<main>
	<h1>Spocker</h1>

	<label for="httpVerbInput">HTTP Verb</label>
	<select id="httpVerbInput" on:change={handleHttpVerb}>
		<option value="GET">GET</option>
		<option value="POST">POST</option>
		<option value="PUT">PUT</option>
		<option value="UPDATE">UPDATE</option>
		<option value="DELETE">DELETE</option>
	</select>

	<label for="endpointInput">Endpoint</label>
	/ <input id="endpointInput" type="text" on:keyup={handleEndpoint} />

	<label for="expectedJsonResponseInput">Expected JSON Response</label>
	<textarea id="expectedJsonResponseInput" type="text" on:keyup={handleExpectedJsonResponse} />

	<br />
	<button id="create" on:click={() => handleCreate()}>Create</button>

	<br />
	<div id="create-message">{createMessage}</div>
</main>

<style>
	main {
		text-align: center;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #0099f6;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 200;
	}

	label {
		margin-bottom: 5px;
	}

	select, input, textarea {
		width: 400px;
		margin-bottom: 30px;
	}

	textarea {
		height: 200px;
	}

	button {
		color: #0099f6;
		background-color: #f2c300;
		text-transform: uppercase;
		border-width: 0px;
		font-weight: 500;
		height: 40px;
		width: 150px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>