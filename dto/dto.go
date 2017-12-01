package dto

// Entity is a named DTO
type Entity struct {
	Name          string
	NameUpperCase string
	Version       int
}

// Activity is a base Entity
var Activity = Entity{Name: "Activity", NameUpperCase: "ACTIVITY", Version: 19}

// ActivitySubType is a base Entity
var ActivitySubType = Entity{Name: "ActivitySubType", NameUpperCase: "ACTIVITYSUBTYPE", Version: 14}

// Approval is a base Entity
var Approval = Entity{Name: "Approval", NameUpperCase: "APPROVAL", Version: 13}

// Attachment is a base Entity
var Attachment = Entity{Name: "Attachment", NameUpperCase: "ATTACHMENT", Version: 16}

// Address is a base Entity
var Address = Entity{Name: "Address", NameUpperCase: "ADDRESS", Version: 17}

// BusinessPartner is a base Entity
var BusinessPartner = Entity{Name: "BusinessPartner", NameUpperCase: "BUSINESSPARTNER", Version: 20}

// BusinessPartnerGroup is a base Entity
var BusinessPartnerGroup = Entity{Name: "BusinessPartnerGroup", NameUpperCase: "BUSINESSPARTNERGROUP", Version: 14}

// BusinessProcessStepDefinition is a base Entity
var BusinessProcessStepDefinition = Entity{Name: "BusinessProcessStepDefinition", NameUpperCase: "BUSINESSPROCESSSTEPDEFINITION", Version: 15}

// ChecklistInstance is a base Entity
var ChecklistInstance = Entity{Name: "ChecklistInstance", NameUpperCase: "CHECKLISTINSTANCE", Version: 17}

// ChecklistCategory is a base Entity
var ChecklistCategory = Entity{Name: "ChecklistCategory", NameUpperCase: "CHECKLISTCATEGORY", Version: 10}

// ChecklistTemplate is a base Entity
var ChecklistTemplate = Entity{Name: "ChecklistTemplate", NameUpperCase: "CHECKLISTTEMPLATE", Version: 17}

// Currency is a base Entity
var Currency = Entity{Name: "Currency", NameUpperCase: "CURRENCY", Version: 11}

// CustomRule is a base Entity
var CustomRule = Entity{Name: "CustomRule", NameUpperCase: "CUSTOMRULE", Version: 8}

// Contact is a base Entity
var Contact = Entity{Name: "Contact", NameUpperCase: "CONTACT", Version: 15}

// CompanyInfo is a base Entity
var CompanyInfo = Entity{Name: "CompanyInfo", NameUpperCase: "COMPANYINFO", Version: 15}

// CompanySettings is a base Entity
var CompanySettings = Entity{Name: "CompanySettings", NameUpperCase: "COMPANYSETTINGS", Version: 11}

// EmployeeBranch is a base Entity
var EmployeeBranch = Entity{Name: "EmployeeBranch", NameUpperCase: "EMPLOYEEBRANCH", Version: 9}

// EmployeeDepartment is a base Entity
var EmployeeDepartment = Entity{Name: "EmployeeDepartment", NameUpperCase: "EMPLOYEEDEPARTMENT", Version: 9}

// EmployeePosition is a base Entity
var EmployeePosition = Entity{Name: "EmployeePosition", NameUpperCase: "EMPLOYEEPOSITION", Version: 9}

// Enumeration is a base Entity
var Enumeration = Entity{Name: "Enumeration", NameUpperCase: "ENUMERATION", Version: 10}

// Equipment is a base Entity
var Equipment = Entity{Name: "Equipment", NameUpperCase: "EQUIPMENT", Version: 16}

// Expense is a base Entity
var Expense = Entity{Name: "Expense", NameUpperCase: "EXPENSE", Version: 15}

// ExpenseType is a base Entity
var ExpenseType = Entity{Name: "ExpenseType", NameUpperCase: "EXPENSETYPE", Version: 15}

// Filter is a base Entity
var Filter = Entity{Name: "Filter", NameUpperCase: "FILTER", Version: 11}

// Function is a base Entity
var Function = Entity{Name: "Function", NameUpperCase: "FUNCTION", Version: 8}

// Group is a base Entity
var Group = Entity{Name: "Group", NameUpperCase: "GROUP", Version: 13}

// Icon is a base Entity
var Icon = Entity{Name: "Icon", NameUpperCase: "ICON", Version: 8}

// Item is a base Entity
var Item = Entity{Name: "Item", NameUpperCase: "ITEM", Version: 21}

// ItemGroup is a base Entity
var ItemGroup = Entity{Name: "ItemGroup", NameUpperCase: "ITEMGROUP", Version: 10}

// ItemPriceListAssignment is a base Entity
var ItemPriceListAssignment = Entity{Name: "ItemPriceListAssignment", NameUpperCase: "ITEMPRICELISTASSIGNMENT", Version: 14}

// Material is a base Entity
var Material = Entity{Name: "Material", NameUpperCase: "MATERIAL", Version: 18}

// Mileage is a base Entity
var Mileage = Entity{Name: "Mileage", NameUpperCase: "MILEAGE", Version: 16}

// MileageType is a base Entity
var MileageType = Entity{Name: "MileageType", NameUpperCase: "MILEAGETYPE", Version: 14}

// PaymentTerm is a base Entity
var PaymentTerm = Entity{Name: "PaymentTerm", NameUpperCase: "PAYMENTTERM", Version: 14}

// Person is a base Entity
var Person = Entity{Name: "Person", NameUpperCase: "PERSON", Version: 18}

// PersonReservation is a base Entity
var PersonReservation = Entity{Name: "PersonReservation", NameUpperCase: "PERSONRESERVATION", Version: 15}

// PersonReservationType is a base Entity
var PersonReservationType = Entity{Name: "PersonReservationType", NameUpperCase: "PERSONRESERVATIONTYPE", Version: 15}

// PersonWorkTimePattern is a base Entity
var PersonWorkTimePattern = Entity{Name: "PersonWorkTimePattern", NameUpperCase: "PERSONWORKTIMEPATTERN", Version: 8}

// Plugin is a base Entity
var Plugin = Entity{Name: "Plugin", NameUpperCase: "PLUGIN", Version: 8}

// Project is a base Entity
var Project = Entity{Name: "Project", NameUpperCase: "PROJECT", Version: 10}

// ProjectPhase is a base Entity
var ProjectPhase = Entity{Name: "ProjectPhase", NameUpperCase: "PROJECTPHASE", Version: 9}

// PriceList is a base Entity
var PriceList = Entity{Name: "PriceList", NameUpperCase: "PRICELIST", Version: 14}

// ProfileObject is a base Entity
var ProfileObject = Entity{Name: "ProfileObject", NameUpperCase: "PROFILEOBJECT", Version: 20}

// ReportTemplate is a base Entity
var ReportTemplate = Entity{Name: "ReportTemplate", NameUpperCase: "REPORTTEMPLATE", Version: 15}

// Requirement is a base Entity
var Requirement = Entity{Name: "Requirement", NameUpperCase: "REQUIREMENT", Version: 8}

// ReservedMaterial is a base Entity
var ReservedMaterial = Entity{Name: "ReservedMaterial", NameUpperCase: "RESERVEDMATERIAL", Version: 16}

// ScreenConfiguration is a base Entity
var ScreenConfiguration = Entity{Name: "ScreenConfiguration", NameUpperCase: "SCREENCONFIGURATION", Version: 8}

// ServiceAssignment is a base Entity
var ServiceAssignment = Entity{Name: "ServiceAssignment", NameUpperCase: "SERVICEASSIGNMENT", Version: 23}

// ServiceAssignmentStatus is a base Entity
var ServiceAssignmentStatus = Entity{Name: "ServiceAssignmentStatus", NameUpperCase: "SERVICEASSIGNMENTSTATUS", Version: 12}

// ServiceAssignmentStatusDefinition is a base Entity
var ServiceAssignmentStatusDefinition = Entity{Name: "ServiceAssignmentStatusDefinition", NameUpperCase: "SERVICEASSIGNMENTSTATUSDEFINITION", Version: 14}

// ServiceCall is a base Entity
var ServiceCall = Entity{Name: "ServiceCall", NameUpperCase: "SERVICECALL", Version: 21}

// ServiceCallProblemType is a base Entity
var ServiceCallProblemType = Entity{Name: "ServiceCallProblemType", NameUpperCase: "SERVICECALLPROBLEMTYPE", Version: 13}

// ServiceCallStatus is a base Entity
var ServiceCallStatus = Entity{Name: "ServiceCallStatus", NameUpperCase: "SERVICECALLSTATUS", Version: 13}

// ServiceCallType is a base Entity
var ServiceCallType = Entity{Name: "ServiceCallType", NameUpperCase: "SERVICECALLTYPE", Version: 12}

// ServiceCallSubject is a base Entity
var ServiceCallSubject = Entity{Name: "ServiceCallSubject", NameUpperCase: "SERVICECALLSUBJECT", Version: 12}

// ServiceCallCode is a base Entity
var ServiceCallCode = Entity{Name: "ServiceCallCode", NameUpperCase: "SERVICECALLCODE", Version: 12}

// ServiceCallResponsible is a base Entity
var ServiceCallResponsible = Entity{Name: "ServiceCallResponsible", NameUpperCase: "SERVICECALLRESPONSIBLE", Version: 12}

// ServiceCallOrigin is a base Entity
var ServiceCallOrigin = Entity{Name: "ServiceCallOrigin", NameUpperCase: "SERVICECALLORIGIN", Version: 13}

// Skill is a base Entity
var Skill = Entity{Name: "Skill", NameUpperCase: "SKILL", Version: 8}

// Team is a base Entity
var Team = Entity{Name: "Team", NameUpperCase: "TEAM", Version: 8}

// Tag is a base Entity
var Tag = Entity{Name: "Tag", NameUpperCase: "TAG", Version: 8}

// Tax is a base Entity
var Tax = Entity{Name: "Tax", NameUpperCase: "TAX", Version: 9}

// TimeEffort is a base Entity
var TimeEffort = Entity{Name: "TimeEffort", NameUpperCase: "TIMEEFFORT", Version: 15}

// TimeTask is a base Entity
var TimeTask = Entity{Name: "TimeTask", NameUpperCase: "TIMETASK", Version: 16}

// TimeSubTask is a base Entity
var TimeSubTask = Entity{Name: "TimeSubTask", NameUpperCase: "TIMESUBTASK", Version: 14}

// Translation is a base Entity
var Translation = Entity{Name: "Translation", NameUpperCase: "TRANSLATION", Version: 8}

// UdfMeta is a base Entity
var UdfMeta = Entity{Name: "UdfMeta", NameUpperCase: "UDFMETA", Version: 13}

// UdfMetaGroup is a base Entity
var UdfMetaGroup = Entity{Name: "UdfMetaGroup", NameUpperCase: "UDFMETAGROUP", Version: 10}

// UserSyncConfirmation is a base Entity
var UserSyncConfirmation = Entity{Name: "UserSyncConfirmation", NameUpperCase: "USERSYNCCONFIRMATION", Version: 13}

// Warehouse is a base Entity
var Warehouse = Entity{Name: "Warehouse", NameUpperCase: "WAREHOUSE", Version: 15}

// WorkTimeTask is a base Entity
var WorkTimeTask = Entity{Name: "WorkTimeTask", NameUpperCase: "WORKTIMETASK", Version: 15}

// WorkTimePattern is a base Entity
var WorkTimePattern = Entity{Name: "WorkTimePattern", NameUpperCase: "WORKTIMEPATTERN", Version: 8}

// WorkTime is a base Entity
var WorkTime = Entity{Name: "WorkTime", NameUpperCase: "WORKTIME", Version: 15}

// CrowdBusinessPartner is a base Entity
var CrowdBusinessPartner = Entity{Name: "CrowdBusinessPartner", NameUpperCase: "CROWDBUSINESSPARTNER", Version: 8}

// CrowdAssignment is a base Entity
var CrowdAssignment = Entity{Name: "CrowdAssignment", NameUpperCase: "CROWDASSIGNMENT", Version: 8}

// Notification is a base Entity
var Notification = Entity{Name: "Notification", NameUpperCase: "NOTIFICATION", Version: 8}

// CrowdExecutionRecord is a base Entity
var CrowdExecutionRecord = Entity{Name: "CrowdExecutionRecord", NameUpperCase: "CROWDEXECUTIONRECORD", Version: 8}

// CrowdPerson is a base Entity
var CrowdPerson = Entity{Name: "CrowdPerson", NameUpperCase: "CROWDPERSON", Version: 8}

// UnifiedPerson is a base Entity
var UnifiedPerson = Entity{Name: "UnifiedPerson", NameUpperCase: "UNIFIEDPERSON", Version: 8}
