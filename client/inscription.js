tipz.controller('tipzHandler', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Inscription";

    $scope.inscriptionClick = function() {
    	$http({
                    method: 'POST',
                    url: "http://localhost:9090/api/createUser",
                    data: {data:
                           { data:
                             {lastname: $scope.InscriptionLastname,
                              firstname: $scope.InscriptionFirstname,
                              mail:  $scope.InscriptionMail,
                              password: $scope.InscriptionPassword,
                             }
                           }
                          },
                    transformRequest: function(obj) {
                        var str = [];
                        for(var p in obj)
                            str.push(encodeURIComponent(p) + "=" + encodeURIComponent(JSON.stringify(obj[p]) ) );
                        return str.join("&");
                    },
                    headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
                }).success(function(data){
                    console.log(data)
                    if (data.success == true) {
						$scope.$parent.$parent.page = "connection.html";

                    }
                })
    }
}]);