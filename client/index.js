var tipz = angular.module('Tipz', []);


tipz.controller('index', ['$scope', '$http', 'NavigationService', function ($scope, $parent, $http, NavigationService) { 
    $scope.page = "'home.html'";
    $scope.isConnected = false;
    $scope.error = "";
}]);


// All data used for navigation
tipz.factory('NavigationService', function() {
    var setProjectID = function (name) {
        this.projectID = name;
    }
    
    var setPage = function (page) {
        this.page = page;
    }
    
    var setReward = function (reward) {
    	this.reward = reward;
    }
    
    var setProject = function (project) {
    	this.project = project;
    }
    var setUser = function (User) {
    	this.User = User;
    }
    var isConnected = function () {
        console.log(this.User)
        return this.User != {};
    }
    
    return {
        projectID: "",
        project: {},
        page : "",
        reward : {},
        User: {},
        setProjectID: setProjectID,
        setPage: setPage,
        setReward: setReward,
        setProject: setProject,
        setUser: setUser,
        isConnected: isConnected
    };
});
