import { findTemplate, instantiateTemplate } from "./templating.js";

// define submit behaviour for adminform
document.getElementById("adminForm").addEventListener("submit", (e) => {
    e.preventDefault();
    submitAdminForm();
    return false;
})

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

export const deleteLink = async (id) => {
    document.getElementById("errorCard").classList.add("hidden");
    const adminPassword = document.querySelector('input[name="adminPassword"]').value;
    const data = {
        id,
        adminPassword,
    };
    try {
        // Send the data to the server
        const response = await fetch('/delete', {
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
        } else if (result.message) {
            document.getElementById(`row-${id}`).remove();
        }
    } catch (error) {
        console.error(error);
        handleError(error);
    }
}

function handleSuccess(links) {
    const adminForm = document.getElementById("adminForm");
    const linksTable = document.getElementById("linksTable");
    const linksList = document.getElementById("linksList");
    adminForm.classList.add("hidden");

    const rowTemplate = findTemplate("rowTemplate");

    for (const link of links) {
        let fullLink = `${window.location.origin}/l/${link.urlEnding}`;
        const fillParams = {
            row: [["id", `row-${link.id}`]],
            id: [["innerText", link.id]],
            clicks: [["innerText", link.clicks]],
            // for whatever reason I couldn't set the onclick of the button directly,
            // so I resort to this instead
            btn: [["innerHTML", `
                    <button onclick="deleteLink(${link.id})" class="pico-background-pink-600">Delete</button>
            `]],
            link: [["innerText", fullLink], ["href", fullLink]],
        }
        const instance = instantiateTemplate(rowTemplate, fillParams);
        linksList.appendChild(instance);
    }
    linksTable.classList.remove("hidden");
}