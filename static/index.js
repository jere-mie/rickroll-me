const submitNewRickroll = async () => {
    // scroll to the top for better visibility of error/success message
    window.scrollTo(0, 0);

    // Get the values from the form inputs
    let urlEnding = document.querySelector('input[name="urlEnding"]').value;
    const siteTitle = document.querySelector('input[name="siteTitle"]').value;
    const siteName = document.querySelector('input[name="siteName"]').value;
    const imgLink = document.querySelector('input[name="imgLink"]').value;
    const siteDescription = document.querySelector('textarea[name="siteDescription"]').value;

    // remove all `/` from url, end encode the result
    urlEnding = encodeURIComponent(urlEnding.replace(/\//g, ''));

    // Create the data object
    const data = {
        urlEnding,
        siteTitle,
        siteName,
        imgLink,
        siteDescription
    };

    console.log(data)


    try {
        // Send the data to the server
        const response = await fetch('/new', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        // Get the response as JSON
        const result = await response.json();

        if (result.error) {
            console.error(result.error);
            handleError(result.error);
        } else if (result.link) {
            handleSuccess(result.link);
        }
    } catch (error) {
        console.error(error);
        handleError(error);
    }
}

function handleError(error) {
    const errorCard = document.getElementById("errorCard");
    const successCard = document.getElementById("successCard");
    successCard.classList.add("hidden");
    errorCard.innerText = error;
    errorCard.classList.remove("hidden");
}

function handleSuccess(link) {
    const errorCard = document.getElementById("errorCard");
    const successCard = document.getElementById("successCard");
    const successLink = document.getElementById("successLink");
    errorCard.classList.add("hidden");

    let fullLink = window.location.origin + link

    successLink.href = fullLink;
    successLink.innerText = fullLink;
    successCard.classList.remove("hidden");
}


  