tipz.controller('connection', ['$scope', '$http', 'UserService', 'NavigationService', function ($scope, $http, UserService, NavigationService) {    
    $scope.title = "Connexion";
    

        $scope.redirection2Inscription = function() {
    	$scope.$parent.$parent.page = "inscription.html";
    };

    $scope.connectionclick = function() {


                $http({
                    method: 'POST',
                    url: "http://localhost:9090/api/connectUser",
                    data: {data:
                           { data:
                             {mail: $scope.connectionName,
                              password: $scope.connectionPassword
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
                    	NavigationService.setUser(data.User);
                    	$scope.$parent.$parent.page = "home.html";
                    }
                    else {
                    	console.log("Wrong Login / Password");
                    }
                });
    };
}]);
