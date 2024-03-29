<script>
	import { onMount } from 'svelte';
	import EndpointCard from './EndpointCard.svelte';

	let baseUrl = 'http://0.0.0.0:5001/';
	let endpoints = [];
	let httpVerbInput = 'GET';
	let endpointInput;
	let expectedJsonResponseInput;
	let createMessageSuccess = '';
	let createMessageFailure = '';
	let endpointsMessage = 'Loading endpoints...';

	onMount(async () => {
		await getAllEndpoints();
	});

	async function getAllEndpoints() {
		await fetch(baseUrl + 'endpoints', {
			method: 'GET',
			mode: 'cors'
		}).catch(_ => {
			endpoints = [];
			endpointsMessage = "Something went wrong. Unable to retrieve endpoints";
		}).then(response => {
			if (response.ok) {
				response.json().then(body => {
					if (body === null) {
						endpoints = [];
						endpointsMessage = "You don't have any endpoints. Create one to get started!";
					} else {
						endpoints = body;
					}
				})
			}
		});
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

		await fetch(baseUrl + 'create', {
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
		<div class="form-row-wrapper">
			<div class="httpVerbInput-wrapper">
				<label class="httpVerbInput-label" for="httpVerbInput">HTTP</label>
				<select class="httpVerbInput-select" id="httpVerbInput" on:change={handleHttpVerb}>
					<option value="GET">GET</option>
					<option value="POST">POST</option>
					<option value="PUT">PUT</option>
					<option value="UPDATE">UPDATE</option>
					<option value="DELETE">DELETE</option>
				</select>
			</div>
	
			<div class="endpointInput-wrapper">
				<label class="endpointInput-label" for="endpointInput">Endpoint</label>
				<div class="endpointInput-input-wrapper">
					<span class="endpointInput-slash">/</span>
					<input class="endpointInput-input" id="endpointInput" type="text" on:keyup={handleEndpoint} />
				</div>
			</div>
		</div>

		<div class="form-row-wrapper">
			<div class="expectedJsonResponseInput-wrapper">
				<label for="expectedJsonResponseInput">Expected JSON Response</label>
				<textarea id="expectedJsonResponseInput" type="text" on:keyup={handleExpectedJsonResponse} />
			</div>
		</div>

		<br />
		<button id="create" on:click={() => handleCreate()}>Create</button>

		<div id="create-message-success">{createMessageSuccess}</div>
		<div id="create-message-failure">{createMessageFailure}</div>
	</div>

	<h1>Endpoints</h1>
	<div id="dashboard">
		{#each endpoints as endpoint}
			<EndpointCard 
				endpoint={endpoint}
				getAll={() => getAllEndpoints()}
			></EndpointCard>
		{/each}
		{#if endpoints.length === 0}
			<div id="endpoints-message">{endpointsMessage}</div>
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
		text-align: left;
		font-size: 12px;
		color: #565656;
	}

	select, input, textarea {
		margin-bottom: 30px;
	}

	textarea {
		height: 200px;
		width: 100%;
	}

	button {
		color: #565656;
		background-color: #f2c300;
		text-transform: uppercase;
		border-width: 0px;
		font-weight: 400;
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

	#endpoints-message {
		color: #565656;
	}

	#create-form {
		margin-bottom: 60px;
	}

	#dashboard {
		margin: 0 auto;
		padding: 0;
		width: 60%;
	}

	#footer {
		padding-bottom: 200px;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}

	.form-row-wrapper {
		margin-left: 25%;
		margin-right: 25%;
		width: 50%;
		display: flex;
		flex-direction: row;
		gap: 12px;
	}

	.httpVerbInput-wrapper {
		width: 15%;
	}

	.httpVerbInput-select, .endpointInput-input, .expectedJsonResponseInput-wrapper {
		width: 100%;
	}

	.endpointInput-wrapper {
		width: 85%;
	}

	.endpointInput-input-wrapper {
		display: flex;
		flex-direction: row;
		gap: 12px;
	}

	.endpointInput-label {
		margin-left: 16px;
	}

	.endpointInput-slash {
		margin-top: 6px;
		color: #565656;
	}
</style>