const submitAdminForm = async () => {
    window.scrollTo(0, 0);
    const adminPassword = document.querySelector('input[name="adminPassword"]').value;
    const data = {
        adminPassword
    }
    try {
        // Send the data to the server
        const response = await fetch('/admin', {
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
        } else if (result.links) {
            handleSuccess(result.links);
        }
    } catch (error) {
        console.error(error);
        handleError(error);
    }
}

function handleError(error) {
    const errorCard = document.getElementById("errorCard");
    errorCard.innerText = error;
    errorCard.classList.remove("hidden");
}

function handleSuccess(links) {
    const adminForm = document.getElementById("adminForm");
    const linksList = document.getElementById("linksList");
    adminForm.classList.add("hidden");

    for(const link of links){
        const li = document.createElement("li");
        li.innerText = `${link.id}: `
        const a = document.createElement("a");
        let fullLink = `${window.location.origin}/l/${link.urlEnding}`;

        a.setAttribute("href", fullLink);
        a.innerText = fullLink;

        li.appendChild(a)
        linksList.appendChild(li)
    }

    linksList.classList.remove("hidden");
    console.table(links);
}