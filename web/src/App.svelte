<script>
	import { onMount } from 'svelte';
	import EndpointCard from './EndpointCard.svelte';

	let endpoints = [];
	let httpVerbInput = "GET";
	let endpointInput;
	let expectedJsonResponseInput;
	let createMessageSuccess = "";
	let createMessageFailure = "";

	onMount(async () => {
		await getAllEndpoints();
	});

	async function getAllEndpoints() {
		await fetch('http://localhost:5005/endpoints', {
			method: 'GET',
			mode: 'cors'
		}).catch(_ => {
			endpoints = [];
		}).then(response => {
			if (response.ok) {
				response.json().then(body => {
					if (body === null) {
						endpoints = [];
					} else {
						endpoints = body;
					}
					console.log(endpoints)
				})
			}
		})
	}

	function handleHttpVerb(e) { httpVerbInput = e.target.value; }
	function handleEndpoint(e) { endpointInput = e.target.value; }
	function handleExpectedJsonResponse(e) { expectedJsonResponseInput = e.target.value; }

	function updateCreateSuccessMessage(newMessage) {
		createMessageSuccess = newMessage;
		createMessageFailure = "";
	}

	function updateCreateFailureMessage(newMessage) {
		createMessageSuccess = "";
		createMessageFailure = newMessage;
	}

	function createFormIsValid() {
		let validations = "";
		if (httpVerbInput === "" || httpVerbInput === undefined) validations += "HTTP Verb Is Empty. ";
		if (endpointInput === "" || endpointInput === undefined) validations += "Endpoint Is Empty. ";
		if (expectedJsonResponseInput === "" || expectedJsonResponseInput === undefined) validations += "Expected JSON Response Is Empty. ";

		if (validations !== "") {
			updateCreateFailureMessage(validations);
			return false;
		}

		return true;
	}

	async function handleCreate() {
		if (!createFormIsValid()) return;

		await fetch('http://localhost:5005/create', {
			method: 'POST',
			mode: 'cors',
			body: JSON.stringify({
				httpVerb: httpVerbInput,
				endpoint: endpointInput,
				expectedJsonResponse: JSON.parse(expectedJsonResponseInput)
			})
		}).catch(_ => {
			updateCreateFailureMessage("Create Failed");
		}).then(response => {
			if (response.ok) {
				updateCreateSuccessMessage("Successfully Created " + httpVerbInput + " " + endpointInput + " Endpoint!");
				getAllEndpoints();
			} else {
				updateCreateFailureMessage("Create Failed");
			}
		})
	}
</script>

<main>
	<h1>Spocker</h1>

	<div id="create-form">
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

		<div id="create-message-success">{createMessageSuccess}</div>
		<div id="create-message-failure">{createMessageFailure}</div>
	</div>

	<h1>Endpoints</h1>
	<div id="dashboard">
		{#each endpoints as endpoint}
			<EndpointCard endpoint={endpoint}></EndpointCard>
		{/each}
		{#if endpoints.length === 0}
			<div>
				You don't have any endpoints. Create one to get started!
			</div>
		{/if}
	</div>

	<div id="footer"></div>
		
</main>

<style>
	main {
		text-align: center;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		margin-top: 20px;
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

	#create-message-success {
		margin-top: 20px;
		color: #0099f6;
	}

	#create-message-failure {
		margin-top: 20px;
		color: red;
	}

	#create-form {
		margin-bottom: 80px;
	}

	#dashboard {
		margin: 0 auto;
		padding: 0;
		width: 1200px;
	}

	#footer {
		padding-bottom: 200px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>