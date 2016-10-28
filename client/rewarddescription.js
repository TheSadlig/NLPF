tipz.controller('rewarddescription', ['$scope', '$http', 'NavigationService', function ($scope, $http, NavigationService) {    
    $scope.title = "Je participe !";
    
    $scope.rewardClick = function() {
			$http({
                    method: 'POST',
                    url: "http://localhost:9090/api/invest",
                    data: {data:
                           { data:
                             {rewardID: NavigationService.reward.ID,
                              userID: NavigationService.User.ID
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
                });
            $scope.$parent.$parent.page = "home.html";
    		};
    $scope.rewardValue = NavigationService.reward.Value;
    $scope.projectName = NavigationService.project.Name;
    $scope.projectAuthor = NavigationService.project.Author;
    $scope.rewardName = NavigationService.reward.Title;
    $scope.rewardValue = NavigationService.reward.Value;
    $scope.rewardDescription = NavigationService.reward.Description;
}]);