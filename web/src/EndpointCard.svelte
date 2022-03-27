<script>
    let endpoint;

    function handleDelete() {}
</script>

<main>
    <div id="endpoint-card-container">
        <div id="endpoint-card">
            {endpoint.id}
            <br />
            {endpoint.httpVerb}
            <br />
            {endpoint.path}
            <br />
            Testing 1 2 3
        </div>
    </div>
</main>

<style>
    #endpoint-card-container {
        display: inline-block;
        margin: 0 auto;
        padding: 0;
        height: 300px;
        width: 33%;
    }

    #endpoint-card {
        margin: 0;
        padding: 10px;
    }
</style>