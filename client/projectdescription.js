tipz.controller('projectdescription', ['$scope', 'UserService', 'NavigationService',
                                       function ($scope, UserService, NavigationService) {    
                                           
    	                                   $scope.title = "";
                                           console.log(NavigationService.projectName);
                                           
                                       }]);
