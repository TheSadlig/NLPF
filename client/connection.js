tipz.controller('connection', ['$scope', '$http', 'UserService', 'NavigationService', function ($scope, $parent, $http, UserService, NavigationService) {    
    $scope.title = "Connexion";
    
    
    $scope.connectionUserID;
    $scope.connectionPassword;

        $scope.redirection2Inscription = function() {
        //console.log($scope.$parent);
    	$scope.$parent.$parent.page = "inscription.html";
    }

    $scope.connectionClick = function() {
    	    $http({
                    method: 'POST',
                    url: "http://localhost:9090/api/connectUser",
                    data: {data:
                           { data:
                             {mail: newProject.name,
                              password: newProject.desc
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
                    if (data.success == false) {
						
                    }
                	if (data.success == true) {
						NavigationService.setUser(data.User);
						
					$scope.$parent.$parent.page = "home.html";

                    }
                });
    }

}]);
