// 示例
var myHeaders = new Headers();
var myInit = {
	method: 'GET',
	headers: myHeaders,
	mode: 'cors',
	cache: 'default'
};
var myRequest = new Request('flower.jpg', myInit);

fetch(myRequest,myInit)
	.then(function(response) {
		return response.blob();
	})
	.then(function(myBlob) {
		var objectURL = URL.createObjectURL(myBlob);
		myImage.src = objectURL;
	});