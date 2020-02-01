package lucy

type DomainType uint

const (
	Where            DomainType  = 0
	Relation         DomainType  = 1
	Creation         DomainType  = 2
	Updation         DomainType  = 3
	Deletion         DomainType  = 4
	SetTarget        DomainType  = 6
	Unknown          DomainType  = 7
	And              DomainType  = 8
	Or               DomainType  = 9
	MiscNodeName     DomainType  = 10 // Made specifically for neo4j
	Model            DomainType  = 11
	ExpressionString DomainType  = 12
	AndStr           DomainType  = 13
	OrStr            DomainType  = 14
	WhereStr         DomainType = 15
)