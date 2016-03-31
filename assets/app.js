function clearAPIRequestAndError() {
    document.getElementById('api_request').value = '';
    document.getElementById('api_request_formatting_error').innerHTML = '';
}

function clearClaimsAndError() {
    document.getElementById('claims').value = '';
    document.getElementById('claims_formatting_error').innerHTML = '';
    console.log('sh')
}

function makeAPIRequest() {
    console.log('makeAPIRequest');
}