package sqlgen

var (
	start              Fn
	switchSysVars      Fn
	createTable        Fn
	colDefs            Fn
	colDef             Fn
	idxDefs            Fn
	idxDef             Fn
	insertInto         Fn
	query              Fn
	commonSelect       Fn
	predicate          Fn
	predicates         Fn
	randColVals        Fn
	forUpdateOpt       Fn
	union              Fn
	aggSelect          Fn
	commonInsert       Fn
	onDupAssignment    Fn
	multipleRowVals    Fn
	multipleRowVal     Fn
	onDuplicateUpdate  Fn
	commonUpdate       Fn
	updateAssignment   Fn
	commonDelete       Fn
	maybeLimit         Fn
	addIndex           Fn
	dropIndex          Fn
	addColumn          Fn
	dropColumn         Fn
	dmlStmt            Fn
	ddlStmt            Fn
	initStart          Fn
	adminCheck         Fn
	randVal            Fn
	cmpSymbol          Fn
	flashBackTable     Fn
	dropTable          Fn
	createTableLike    Fn
	multiTableQuery    Fn
	joinPredicates     Fn
	joinPredicate      Fn
	splitRegion        Fn
	partitionDef       Fn
	selectIntoOutFile  Fn
	loadTable          Fn
	commonAnalyze      Fn
	prepareStmt        Fn
	queryPrepare       Fn
	deallocPrepareStmt Fn
)
