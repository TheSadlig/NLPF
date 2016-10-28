tipz.controller('projectdescription', ['$scope', '$http', 'UserService', 'NavigationService',
                                       function ($scope, $http, UserService, NavigationService) {    
                
                
                 
    	    
    	    $scope.loadFeed = function(e, p) {
            NavigationService.setReward(p);
            NavigationService.setProject($scope.project);
            console.log($scope.project);
            $scope.$parent.$parent.page = "rewarddescription.html";
    		};
                          
                $http({
                    method: 'POST',
                    url: "http://localhost:9090/api/getProjectById",
                    data: {data:
                           { data:
                             {ID: NavigationService.projectID
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
                    //console.log(data);
                    $scope.project = data.data;
                    $scope.projectDescription = data.data.Description;
                    $scope.projectAuthor = data.data.Author;
                    $scope.projectMail = data.data.Mail;
                    $scope.projectRewards = data.data.Rewards;
                    $scope.title = data.data.Name;
                });
    }]);
