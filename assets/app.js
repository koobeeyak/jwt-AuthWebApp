function clearAPIRequestAndError() {
    document.getElementById('api_request').value = '';
    document.getElementById('api_request_formatting_error').innerHTML = '';
}

function clearClaimsAndError() {
    document.getElementById('claims').value = '';
    document.getElementById('claims_formatting_error').innerHTML = '';
}

function makeAPIRequest() {
    console.log('makeAPIRequest');
    $.ajax({
        url: "/getjwt",
        type: "POST",
        data: {path: $("#url").val(), claims: $("#claims").val()}
    }).error(function(e, x, y) {
        console.log("Did not get jwt, jqXHR:", e, x, y);
    }).done(function(response, status, jqXHR) {
        console.log("Success?",response, status, jqXHR);
    });
}
