app;
app.controller('mainController', ['$scope', '$http', function($scope, $http) {

//	$scope.items = [
//	"The first choice!",
//	"And another choice for you.",
//	"but wait! A third!"
//	];
	$scope.signin = false;
	$scope.signup = true;
	$scope.formToggle = function() {
		$scope.signin = !$scope.signin;
		$scope.signup = !$scope.signup;
		console.log("signin " + $scope.signin)
		console.log("signup " + $scope.signup)
	}

	$scope.processForm = function() {
		if ($scope.signupForm.$valid == true) {
			alert("checkformemail " + $scope.suForm.email); //su = Sign Up
			console.log("valid " + $scope.signupForm.$valid); 

			$http({
				method  : 'POST',
				url     : '/signup',
				data    : $.param($scope.suForm),  // pass in data as strings, su = Sign Up
				headers : { 'Content-Type': 'application/x-www-form-urlencoded' }  // set the headers so angular passing info as form data (not request payload)
			})

		};
	};

	$scope.loginForm = function() {
		//alert("checkusernameinput" + $scope.siForm.username);
		//if ($scope.signinForm.$valid == true) {
		alert("checkusernameinput" + $scope.siForm.suUsername);
		$http({
			method  : 'POST',
			url     : '/signin',
			data    : $.param($scope.siForm),  // pass in data as strings, si = Sign In
			headers : { 'Content-Type': 'application/x-www-form-urlencoded' }  // set the headers so angular passing info as form data (not request payload)
		})
		//};
	};
	$scope.pw1 = '';

}]);

app
.directive('pwCheck', [function () {
	return {
		require: 'ngModel',
		link: function (scope, elem, attrs, ctrl) {
			var firstPassword = '#' + attrs.pwCheck;
			elem.add(firstPassword).on('keyup', function () {
				scope.$apply(function () {
					ctrl.$setValidity('pwmatch', elem.val() === $(firstPassword).val());
				});
			});
		}
	}
}]);

