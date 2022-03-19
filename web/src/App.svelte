<script>
	let httpVerb;
	let endpoint;
	let expectedJsonResponse;
</script>

<main>
	<h1>Spocker</h1>

	<label for="httpVerbInput">HTTP Verb</label>
	<select id="httpVerbInput">
		<option value="get">GET</option>
		<option value="post">POST</option>
		<option value="put">PUT</option>
		<option value="update">UPDATE</option>
		<option value="delete">DELETE</option>
	</select>

	<label for="endpointInput">Endpoint</label>
	<input id="endpointInput" type="text" />

	<label for="expectedJsonResponseInput">Expected JSON Response</label>
	<textarea id="expectedJsonResponseInput" type="text" />

	<br />
	<button>Create</button>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
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

	select, input {
		width: 400px;
		margin-bottom: 30px;
	}

	textarea {
		height: 200px;
		width: 400px;
		margin-bottom: 30px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>