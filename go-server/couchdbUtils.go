package main
import ("fmt"
	"github.com/rhinoman/couchdb-go"
	"time"
	"github.com/nu7hatch/gouuid"
)

type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []interface{} `json:"rows,omitempty"`
}

type MultiReadResponse struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Rows      []MultiReadRow `json:"rows"`
}

type MultiReadRow struct {
	Id  string       `json:"id"`
	Key string       `json:"key"`
	Doc interface{} `json:"doc"`
}

type ListResult struct {
	Id  string       `json:"id"`
	Key interface{} `json:"key"`
	//Value string       `json:"value"`
}

type ListResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ListResult `json:"rows,omitempty"`
}

type View struct {
	Map    string `json:"map"`
	Reduce string `json:"reduce,omitempty"`
}

type DesignDocument struct {
	Language string            `json:"language"`
	Views    map[string]View   `json:"views"`
	Lists    map[string]string `json:"lists"`
}

func getAuth() couchdb.BasicAuth {
	return couchdb.BasicAuth{"admin", "admin"}
}

func getConn () *couchdb.Connection {
	var timeout = time.Duration(50000 * time.Millisecond)
	conn, _ := couchdb.NewConnection("127.0.0.1",5984,timeout)
	return conn

}
func createDB () {
	conn := getConn()
	auth := getAuth()
	conn.CreateDB("nlpf", &auth);
}

func getDB () *couchdb.Database {
	conn := getConn()
	auth := getAuth()
	db := conn.SelectDB("nlpf", &auth)
	
	return db
}

func getUUID() string {
	out, _ := uuid.NewV4()
	return out.String()
}

func createDummyProject() string{
	db := getDB()

	u := getEmptyUser()
	u.Lastname = "Thiaw-Kine"
	u.Firstname = "Roman"
	u.Mail = "thiaw-_r@epita.fr"
	u.Password = "dicks"

	_, err := db.Save(&u, u.ID, "")
	
	project := getEmptyProject()
	
	project.Name= "ProjName"
	project.Description= "DescProj"
	project.User_ID= u.ID
	project.Date = "03/12/1993"
	
	_, err = db.Save(&project, project.ID, "")

	r:= getEmptyReward()	
	r.Title = "Pas cher"
	r.Description = "c'est trop pas cher"
	r.Value = 500
	r.Project_ID = project.ID

	r2:= getEmptyReward()	
	r2.Title = "Pas cher2"
	r2.Description = "c'est trop pas cher2"
	r2.Value = 2500
	r2.Project_ID = project.ID
	
	r3:= getEmptyReward()	
	r3.Title = "Pas cher3"
	r3.Description = "c'est trop pas cher3"
	r3.Value = 3500
	r3.Project_ID = project.ID
	

	id, err := db.Save(&r, r.ID, "")
	id, err = db.Save(&r2, r2.ID, "")
	id, err = db.Save(&r3, r3.ID, "")

	emptyProject := Project{}
	_, err = db.Read(project.ID, &emptyProject, nil)
	fmt.Println(emptyProject)
	//	u := User{Lastname: "Gildas", Firstname: "Lebel", Mail: "ezmfiej@qzefq.fe", Password: "ezeq"}


	
	fmt.Println("err", err)
	fmt.Println("DocRev", id)
	return project.ID
}

func getUserById(userID string) *User {
	db := getDB();
	
	result := User{}
	_, _ = db.Read(userID, &result, nil)

	return &result
}





func createView(viewName string, designName string, mapCode string, reduceCode string) {
	db := getDB();
	
	
	view := View{
		Map: mapCode}
	if (reduceCode != "") {
		view.Reduce = reduceCode
	}
		
	views := make(map[string]View)
	
	views[viewName] = view
	ddoc := DesignDocument{
		Language: "javascript",
		Views:    views}
	
	db.SaveDesignDoc(designName, ddoc, "")


}
func createViews() {
	createDummyProject()

	createView("get_projects", "projects", "function(doc) {\n  if (doc.Type == \"Project\"){\n    emit(doc.ID, doc)\n  }\n}", "")
	createView("get_rewards", "rewards", "function(doc) {\n if (doc.Type == \"Reward\") {\n    emit(doc.Project_ID, doc);\n  }\n}", "")
	createView("get_users_by_mail", "user_mail", "function(doc) {\n if (doc.Type == \"User\") {\n    emit(doc.Mail, doc);\n  }\n}", "")
	createView("get_users", "user", "function(doc) {\n if (doc.Type == \"User\") {\n    emit(doc.ID, doc);\n  }\n}", "")
	createView("get_investments_number_by_reward",
		"investment",
		"function(doc) {\n if (doc.Type == \"Investment\")\n emit(doc.Reward_ID, 1);\n }",
		"function(keys, values) { \n return sum(values); \n }")
	
}
