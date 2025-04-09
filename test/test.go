package main

import (
	"fmt"
	"github.com/892294101/sqlbuilder"
)

func main() {

	/*var addUserOpt []sqlbuilder.InsertOption
	addUserOpt = append(addUserOpt, sqlbuilder.SetInsertColumn("role_name"))
	addUserOpt = append(addUserOpt, sqlbuilder.SetInsertColumn("role_name"))
	addUserOpt = append(addUserOpt, sqlbuilder.SetInsertColumn("role_name"))
	addUserOpt = append(addUserOpt, sqlbuilder.SetInsertColumn("role_name"))
	addUserOpt = append(addUserOpt, sqlbuilder.SetInsertColumn("role_name"))
	t, e := sqlbuilder.SQLB.GetInsert(sqlbuilder.InternalInsert, addUserOpt...)
	fmt.Println(t, e)

	var ss []sqlbuilder.DeleteOption
	ss = append(ss, sqlbuilder.SetDeleteWhereColumn("and", "name", "="))
	ss = append(ss, sqlbuilder.SetDeleteWhereColumn("and", "name", "="))
	ss = append(ss, sqlbuilder.SetDeleteWhereColumn("and", "name", "="))
	ss = append(ss, sqlbuilder.SetDeleteWhereColumn("and", "name", "="))
	ss = append(ss, sqlbuilder.SetDeleteWhereColumn("and", "name", "="))
	t, e = sqlbuilder.SQLB.GetDelete(sqlbuilder.InternalDelete, ss...)
	fmt.Println(t, e)

	var sss []sqlbuilder.UpdateOption
	sss = append(sss, sqlbuilder.SetUpdateSet("name"))
	sss = append(sss, sqlbuilder.SetUpdateSet("sex1"))
	sss = append(sss, sqlbuilder.SetUpdateSet("sex2"))
	sss = append(sss, sqlbuilder.SetUpdateSet("sex3"))
	sss = append(sss, sqlbuilder.SetUpdateSet("sex4"))
	sss = append(sss, sqlbuilder.SetUpdateWhereColumn("and", "name", "="))
	sss = append(sss, sqlbuilder.SetUpdateWhereColumn("and", "name", "="))
	sss = append(sss, sqlbuilder.SetUpdateWhereColumn("and", "name", "="))
	sss = append(sss, sqlbuilder.SetUpdateWhereColumn("and", "name", "="))
	t, e = sqlbuilder.SQLB.GetUpdate(sqlbuilder.InternalUpdate, sss...)

	fmt.Println(t, e)

	var ssss []sqlbuilder.SelectOption
	ssss = append(ssss, sqlbuilder.SetSelectResultColumn("id"))
	ssss = append(ssss, sqlbuilder.SetSelectWhereColumn("and", "name", "="))
	ssss = append(ssss, sqlbuilder.SetSelectWhereColumn("and", "address", "="))
	ssss = append(ssss, sqlbuilder.SetSelectWhereColumn("and", "address", "="))
	ssss = append(ssss, sqlbuilder.SetSelectWhereColumn("and", "address", "="))

	sqlbuilder.SQLB.SetDBType("oracle")
	t, e = sqlbuilder.SQLB.GetSelect(sqlbuilder.InternalSelect, ssss...)

	fmt.Println(t, e)

	var cs []sqlbuilder.CountOption
	cs = append(cs, sqlbuilder.SetCountWhereColumn("and", "s1", "="))
	cs = append(cs, sqlbuilder.SetCountWhereColumn("and", "s2", "="))
	cs = append(cs, sqlbuilder.SetCountWhereColumn("and", "s3", "="))
	cs = append(cs, sqlbuilder.SetCountWhereColumn("and", "s4", "="))
	sqlbuilder.SQLB.SetDBType("oracle")
	t, e = sqlbuilder.SQLB.GetCount(sqlbuilder.InternalSelect, cs...)
	fmt.Println("----------------------------- ", t, e)

	///////////////////////////////////////////////////////////////////

	t, e = sqlbuilder.SQLB.GetPublicSelect("lender_new01", ssss...)
	fmt.Println(t, e)

	t, e = sqlbuilder.SQLB.GetPublicInsert("lender_new01", addUserOpt...)
	fmt.Println(t, e)

	t, e = sqlbuilder.SQLB.GetPublicUpdate("lender_new01", sss...)
	fmt.Println(t, e)

	t, e = sqlbuilder.SQLB.GetPublicDelete("lender_new01", ss...)

	fmt.Println(t, e)
	*/
	var cs1 []sqlbuilder.SelectOption
	cs1 = append(cs1, sqlbuilder.SetSelectResultColumn("id"))

	cs1 = append(cs1, sqlbuilder.SetSelectWhereColumn("and", "aaaaaaaaaaaaaaaaaa", "="))
	cs1 = append(cs1, sqlbuilder.SetSelectWhereColumnIsNull("and", "asdfsadfsgdsfgdsfg"))
	cs1 = append(cs1, sqlbuilder.SetSelectWhereColumn("and", "s2", "="))
	cs1 = append(cs1, sqlbuilder.SetSelectWhereColumnNotNull("or", "asdfsadfsa"))
	sqlbuilder.SQLB.SetDBType("oracle")
	t, e := sqlbuilder.SQLB.GetSelect(sqlbuilder.InternalSelect, cs1...)
	fmt.Println("====================", t, e)

}
