//async function upload(formData) {
//  try {
//    const response = await fetch("http://35.212.179.245/", {
//      method: "PUT",
//      body: formData,
//    });
//    const result = await response.json();
//    console.log("Success:", result);
//  } catch (error) {
//    console.error("Error:", error);
//  }
//}
//
//const formData = new FormData();
//const fileField = document.querySelector('input[type="file"]');
//
//formData.append("username", "abc123");
//formData.append("avatar", fileField.files[0]);
//
//upload(formData);

const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit);

async function handleSubmit(event) {
  /** @type {HTMLFormElement} */
	const form = event.currentTarget;
	const url = new URL(form.action);
	const formData = new FormData(form);
	const searchParams = new URLSearchParams(formData);

	/** @type {Parameters<fetch>[1]} */
	const fetchOptions = {
		method: form.method,
	};

	if (form.method.toLowerCase() === 'post') {
	if (form.enctype === 'multipart/form-data') {
		fetchOptions.body = formData;
	} else {
		fetchOptions.body = searchParams;
	}
	} else {
		url.search = searchParams;
	}

	fetch(url, fetchOptions);

	event.preventDefault();
}
