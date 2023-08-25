// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ContactType string

// Enum values for ContactType
const (
	ContactTypePerson      ContactType = "PERSON"
	ContactTypeCompany     ContactType = "COMPANY"
	ContactTypeAssociation ContactType = "ASSOCIATION"
	ContactTypePublicBody  ContactType = "PUBLIC_BODY"
	ContactTypeReseller    ContactType = "RESELLER"
)

// Values returns all known values for ContactType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ContactType) Values() []ContactType {
	return []ContactType{
		"PERSON",
		"COMPANY",
		"ASSOCIATION",
		"PUBLIC_BODY",
		"RESELLER",
	}
}

type CountryCode string

// Enum values for CountryCode
const (
	CountryCodeAc CountryCode = "AC"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAe CountryCode = "AE"
	CountryCodeAf CountryCode = "AF"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAl CountryCode = "AL"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAn CountryCode = "AN"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBi CountryCode = "BI"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBl CountryCode = "BL"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBr CountryCode = "BR"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCf CountryCode = "CF"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCh CountryCode = "CH"
	CountryCodeCi CountryCode = "CI"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCo CountryCode = "CO"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCv CountryCode = "CV"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDe CountryCode = "DE"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEe CountryCode = "EE"
	CountryCodeEg CountryCode = "EG"
	CountryCodeEh CountryCode = "EH"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEs CountryCode = "ES"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFm CountryCode = "FM"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGb CountryCode = "GB"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGe CountryCode = "GE"
	CountryCodeGf CountryCode = "GF"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGs CountryCode = "GS"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHm CountryCode = "HM"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHr CountryCode = "HR"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHu CountryCode = "HU"
	CountryCodeId CountryCode = "ID"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIn CountryCode = "IN"
	CountryCodeIo CountryCode = "IO"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJo CountryCode = "JO"
	CountryCodeJp CountryCode = "JP"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKg CountryCode = "KG"
	CountryCodeKh CountryCode = "KH"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKm CountryCode = "KM"
	CountryCodeKn CountryCode = "KN"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKy CountryCode = "KY"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLc CountryCode = "LC"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLk CountryCode = "LK"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLy CountryCode = "LY"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMf CountryCode = "MF"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMm CountryCode = "MM"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMp CountryCode = "MP"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMu CountryCode = "MU"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMx CountryCode = "MX"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNf CountryCode = "NF"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNo CountryCode = "NO"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeOm CountryCode = "OM"
	CountryCodePa CountryCode = "PA"
	CountryCodePe CountryCode = "PE"
	CountryCodePf CountryCode = "PF"
	CountryCodePg CountryCode = "PG"
	CountryCodePh CountryCode = "PH"
	CountryCodePk CountryCode = "PK"
	CountryCodePl CountryCode = "PL"
	CountryCodePm CountryCode = "PM"
	CountryCodePn CountryCode = "PN"
	CountryCodePr CountryCode = "PR"
	CountryCodePs CountryCode = "PS"
	CountryCodePt CountryCode = "PT"
	CountryCodePw CountryCode = "PW"
	CountryCodePy CountryCode = "PY"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRs CountryCode = "RS"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSe CountryCode = "SE"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSh CountryCode = "SH"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSn CountryCode = "SN"
	CountryCodeSo CountryCode = "SO"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSs CountryCode = "SS"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSv CountryCode = "SV"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSy CountryCode = "SY"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTd CountryCode = "TD"
	CountryCodeTf CountryCode = "TF"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTp CountryCode = "TP"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTv CountryCode = "TV"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeUa CountryCode = "UA"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUs CountryCode = "US"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVa CountryCode = "VA"
	CountryCodeVc CountryCode = "VC"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVu CountryCode = "VU"
	CountryCodeWf CountryCode = "WF"
	CountryCodeWs CountryCode = "WS"
	CountryCodeYe CountryCode = "YE"
	CountryCodeYt CountryCode = "YT"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

// Values returns all known values for CountryCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CountryCode) Values() []CountryCode {
	return []CountryCode{
		"AC",
		"AD",
		"AE",
		"AF",
		"AG",
		"AI",
		"AL",
		"AM",
		"AN",
		"AO",
		"AQ",
		"AR",
		"AS",
		"AT",
		"AU",
		"AW",
		"AX",
		"AZ",
		"BA",
		"BB",
		"BD",
		"BE",
		"BF",
		"BG",
		"BH",
		"BI",
		"BJ",
		"BL",
		"BM",
		"BN",
		"BO",
		"BQ",
		"BR",
		"BS",
		"BT",
		"BV",
		"BW",
		"BY",
		"BZ",
		"CA",
		"CC",
		"CD",
		"CF",
		"CG",
		"CH",
		"CI",
		"CK",
		"CL",
		"CM",
		"CN",
		"CO",
		"CR",
		"CU",
		"CV",
		"CW",
		"CX",
		"CY",
		"CZ",
		"DE",
		"DJ",
		"DK",
		"DM",
		"DO",
		"DZ",
		"EC",
		"EE",
		"EG",
		"EH",
		"ER",
		"ES",
		"ET",
		"FI",
		"FJ",
		"FK",
		"FM",
		"FO",
		"FR",
		"GA",
		"GB",
		"GD",
		"GE",
		"GF",
		"GG",
		"GH",
		"GI",
		"GL",
		"GM",
		"GN",
		"GP",
		"GQ",
		"GR",
		"GS",
		"GT",
		"GU",
		"GW",
		"GY",
		"HK",
		"HM",
		"HN",
		"HR",
		"HT",
		"HU",
		"ID",
		"IE",
		"IL",
		"IM",
		"IN",
		"IO",
		"IQ",
		"IR",
		"IS",
		"IT",
		"JE",
		"JM",
		"JO",
		"JP",
		"KE",
		"KG",
		"KH",
		"KI",
		"KM",
		"KN",
		"KP",
		"KR",
		"KW",
		"KY",
		"KZ",
		"LA",
		"LB",
		"LC",
		"LI",
		"LK",
		"LR",
		"LS",
		"LT",
		"LU",
		"LV",
		"LY",
		"MA",
		"MC",
		"MD",
		"ME",
		"MF",
		"MG",
		"MH",
		"MK",
		"ML",
		"MM",
		"MN",
		"MO",
		"MP",
		"MQ",
		"MR",
		"MS",
		"MT",
		"MU",
		"MV",
		"MW",
		"MX",
		"MY",
		"MZ",
		"NA",
		"NC",
		"NE",
		"NF",
		"NG",
		"NI",
		"NL",
		"NO",
		"NP",
		"NR",
		"NU",
		"NZ",
		"OM",
		"PA",
		"PE",
		"PF",
		"PG",
		"PH",
		"PK",
		"PL",
		"PM",
		"PN",
		"PR",
		"PS",
		"PT",
		"PW",
		"PY",
		"QA",
		"RE",
		"RO",
		"RS",
		"RU",
		"RW",
		"SA",
		"SB",
		"SC",
		"SD",
		"SE",
		"SG",
		"SH",
		"SI",
		"SJ",
		"SK",
		"SL",
		"SM",
		"SN",
		"SO",
		"SR",
		"SS",
		"ST",
		"SV",
		"SX",
		"SY",
		"SZ",
		"TC",
		"TD",
		"TF",
		"TG",
		"TH",
		"TJ",
		"TK",
		"TL",
		"TM",
		"TN",
		"TO",
		"TP",
		"TR",
		"TT",
		"TV",
		"TW",
		"TZ",
		"UA",
		"UG",
		"US",
		"UY",
		"UZ",
		"VA",
		"VC",
		"VE",
		"VG",
		"VI",
		"VN",
		"VU",
		"WF",
		"WS",
		"YE",
		"YT",
		"ZA",
		"ZM",
		"ZW",
	}
}

type DomainAvailability string

// Enum values for DomainAvailability
const (
	DomainAvailabilityAvailable             DomainAvailability = "AVAILABLE"
	DomainAvailabilityAvailableReserved     DomainAvailability = "AVAILABLE_RESERVED"
	DomainAvailabilityAvailablePreorder     DomainAvailability = "AVAILABLE_PREORDER"
	DomainAvailabilityUnavailable           DomainAvailability = "UNAVAILABLE"
	DomainAvailabilityUnavailablePremium    DomainAvailability = "UNAVAILABLE_PREMIUM"
	DomainAvailabilityUnavailableRestricted DomainAvailability = "UNAVAILABLE_RESTRICTED"
	DomainAvailabilityReserved              DomainAvailability = "RESERVED"
	DomainAvailabilityDontKnow              DomainAvailability = "DONT_KNOW"
)

// Values returns all known values for DomainAvailability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainAvailability) Values() []DomainAvailability {
	return []DomainAvailability{
		"AVAILABLE",
		"AVAILABLE_RESERVED",
		"AVAILABLE_PREORDER",
		"UNAVAILABLE",
		"UNAVAILABLE_PREMIUM",
		"UNAVAILABLE_RESTRICTED",
		"RESERVED",
		"DONT_KNOW",
	}
}

type ExtraParamName string

// Enum values for ExtraParamName
const (
	ExtraParamNameDunsNumber                    ExtraParamName = "DUNS_NUMBER"
	ExtraParamNameBrandNumber                   ExtraParamName = "BRAND_NUMBER"
	ExtraParamNameBirthDepartment               ExtraParamName = "BIRTH_DEPARTMENT"
	ExtraParamNameBirthDateInYyyyMmDd           ExtraParamName = "BIRTH_DATE_IN_YYYY_MM_DD"
	ExtraParamNameBirthCountry                  ExtraParamName = "BIRTH_COUNTRY"
	ExtraParamNameBirthCity                     ExtraParamName = "BIRTH_CITY"
	ExtraParamNameDocumentNumber                ExtraParamName = "DOCUMENT_NUMBER"
	ExtraParamNameAuIdNumber                    ExtraParamName = "AU_ID_NUMBER"
	ExtraParamNameAuIdType                      ExtraParamName = "AU_ID_TYPE"
	ExtraParamNameCaLegalType                   ExtraParamName = "CA_LEGAL_TYPE"
	ExtraParamNameCaBusinessEntityType          ExtraParamName = "CA_BUSINESS_ENTITY_TYPE"
	ExtraParamNameCaLegalRepresentative         ExtraParamName = "CA_LEGAL_REPRESENTATIVE"
	ExtraParamNameCaLegalRepresentativeCapacity ExtraParamName = "CA_LEGAL_REPRESENTATIVE_CAPACITY"
	ExtraParamNameEsIdentification              ExtraParamName = "ES_IDENTIFICATION"
	ExtraParamNameEsIdentificationType          ExtraParamName = "ES_IDENTIFICATION_TYPE"
	ExtraParamNameEsLegalForm                   ExtraParamName = "ES_LEGAL_FORM"
	ExtraParamNameFiBusinessNumber              ExtraParamName = "FI_BUSINESS_NUMBER"
	ExtraParamNameOnwerFiIdNumber               ExtraParamName = "FI_ID_NUMBER"
	ExtraParamNameFiNationality                 ExtraParamName = "FI_NATIONALITY"
	ExtraParamNameFiOrganizationType            ExtraParamName = "FI_ORGANIZATION_TYPE"
	ExtraParamNameItNationality                 ExtraParamName = "IT_NATIONALITY"
	ExtraParamNameItPin                         ExtraParamName = "IT_PIN"
	ExtraParamNameItRegistrantEntityType        ExtraParamName = "IT_REGISTRANT_ENTITY_TYPE"
	ExtraParamNameRuPassportData                ExtraParamName = "RU_PASSPORT_DATA"
	ExtraParamNameSeIdNumber                    ExtraParamName = "SE_ID_NUMBER"
	ExtraParamNameSgIdNumber                    ExtraParamName = "SG_ID_NUMBER"
	ExtraParamNameVatNumber                     ExtraParamName = "VAT_NUMBER"
	ExtraParamNameUkContactType                 ExtraParamName = "UK_CONTACT_TYPE"
	ExtraParamNameUkCompanyNumber               ExtraParamName = "UK_COMPANY_NUMBER"
	ExtraParamNameEuCountryOfCitizenship        ExtraParamName = "EU_COUNTRY_OF_CITIZENSHIP"
	ExtraParamNameAuPriorityToken               ExtraParamName = "AU_PRIORITY_TOKEN"
)

// Values returns all known values for ExtraParamName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExtraParamName) Values() []ExtraParamName {
	return []ExtraParamName{
		"DUNS_NUMBER",
		"BRAND_NUMBER",
		"BIRTH_DEPARTMENT",
		"BIRTH_DATE_IN_YYYY_MM_DD",
		"BIRTH_COUNTRY",
		"BIRTH_CITY",
		"DOCUMENT_NUMBER",
		"AU_ID_NUMBER",
		"AU_ID_TYPE",
		"CA_LEGAL_TYPE",
		"CA_BUSINESS_ENTITY_TYPE",
		"CA_LEGAL_REPRESENTATIVE",
		"CA_LEGAL_REPRESENTATIVE_CAPACITY",
		"ES_IDENTIFICATION",
		"ES_IDENTIFICATION_TYPE",
		"ES_LEGAL_FORM",
		"FI_BUSINESS_NUMBER",
		"FI_ID_NUMBER",
		"FI_NATIONALITY",
		"FI_ORGANIZATION_TYPE",
		"IT_NATIONALITY",
		"IT_PIN",
		"IT_REGISTRANT_ENTITY_TYPE",
		"RU_PASSPORT_DATA",
		"SE_ID_NUMBER",
		"SG_ID_NUMBER",
		"VAT_NUMBER",
		"UK_CONTACT_TYPE",
		"UK_COMPANY_NUMBER",
		"EU_COUNTRY_OF_CITIZENSHIP",
		"AU_PRIORITY_TOKEN",
	}
}

type ListDomainsAttributeName string

// Enum values for ListDomainsAttributeName
const (
	ListDomainsAttributeNameDomainName ListDomainsAttributeName = "DomainName"
	ListDomainsAttributeNameExpiry     ListDomainsAttributeName = "Expiry"
)

// Values returns all known values for ListDomainsAttributeName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListDomainsAttributeName) Values() []ListDomainsAttributeName {
	return []ListDomainsAttributeName{
		"DomainName",
		"Expiry",
	}
}

type ListOperationsSortAttributeName string

// Enum values for ListOperationsSortAttributeName
const (
	ListOperationsSortAttributeNameSubmittedDate ListOperationsSortAttributeName = "SubmittedDate"
)

// Values returns all known values for ListOperationsSortAttributeName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ListOperationsSortAttributeName) Values() []ListOperationsSortAttributeName {
	return []ListOperationsSortAttributeName{
		"SubmittedDate",
	}
}

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusSubmitted  OperationStatus = "SUBMITTED"
	OperationStatusInProgress OperationStatus = "IN_PROGRESS"
	OperationStatusError      OperationStatus = "ERROR"
	OperationStatusSuccessful OperationStatus = "SUCCESSFUL"
	OperationStatusFailed     OperationStatus = "FAILED"
)

// Values returns all known values for OperationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationStatus) Values() []OperationStatus {
	return []OperationStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"ERROR",
		"SUCCESSFUL",
		"FAILED",
	}
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeRegisterDomain            OperationType = "REGISTER_DOMAIN"
	OperationTypeDeleteDomain              OperationType = "DELETE_DOMAIN"
	OperationTypeTransferInDomain          OperationType = "TRANSFER_IN_DOMAIN"
	OperationTypeUpdateDomainContact       OperationType = "UPDATE_DOMAIN_CONTACT"
	OperationTypeUpdateNameserver          OperationType = "UPDATE_NAMESERVER"
	OperationTypeChangePrivacyProtection   OperationType = "CHANGE_PRIVACY_PROTECTION"
	OperationTypeDomainLock                OperationType = "DOMAIN_LOCK"
	OperationTypeEnableAutorenew           OperationType = "ENABLE_AUTORENEW"
	OperationTypeDisableAutorenew          OperationType = "DISABLE_AUTORENEW"
	OperationTypeAddDnssec                 OperationType = "ADD_DNSSEC"
	OperationTypeRemoveDnssec              OperationType = "REMOVE_DNSSEC"
	OperationTypeExpireDomain              OperationType = "EXPIRE_DOMAIN"
	OperationTypeTransferOutDomain         OperationType = "TRANSFER_OUT_DOMAIN"
	OperationTypeChangeDomainOwner         OperationType = "CHANGE_DOMAIN_OWNER"
	OperationTypeRenewDomain               OperationType = "RENEW_DOMAIN"
	OperationTypePushDomain                OperationType = "PUSH_DOMAIN"
	OperationTypeInternalTransferOutDomain OperationType = "INTERNAL_TRANSFER_OUT_DOMAIN"
	OperationTypeInternalTransferInDomain  OperationType = "INTERNAL_TRANSFER_IN_DOMAIN"
)

// Values returns all known values for OperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationType) Values() []OperationType {
	return []OperationType{
		"REGISTER_DOMAIN",
		"DELETE_DOMAIN",
		"TRANSFER_IN_DOMAIN",
		"UPDATE_DOMAIN_CONTACT",
		"UPDATE_NAMESERVER",
		"CHANGE_PRIVACY_PROTECTION",
		"DOMAIN_LOCK",
		"ENABLE_AUTORENEW",
		"DISABLE_AUTORENEW",
		"ADD_DNSSEC",
		"REMOVE_DNSSEC",
		"EXPIRE_DOMAIN",
		"TRANSFER_OUT_DOMAIN",
		"CHANGE_DOMAIN_OWNER",
		"RENEW_DOMAIN",
		"PUSH_DOMAIN",
		"INTERNAL_TRANSFER_OUT_DOMAIN",
		"INTERNAL_TRANSFER_IN_DOMAIN",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorLe         Operator = "LE"
	OperatorGe         Operator = "GE"
	OperatorBeginsWith Operator = "BEGINS_WITH"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"LE",
		"GE",
		"BEGINS_WITH",
	}
}

type ReachabilityStatus string

// Enum values for ReachabilityStatus
const (
	ReachabilityStatusPending ReachabilityStatus = "PENDING"
	ReachabilityStatusDone    ReachabilityStatus = "DONE"
	ReachabilityStatusExpired ReachabilityStatus = "EXPIRED"
)

// Values returns all known values for ReachabilityStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReachabilityStatus) Values() []ReachabilityStatus {
	return []ReachabilityStatus{
		"PENDING",
		"DONE",
		"EXPIRED",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASC",
		"DESC",
	}
}

type StatusFlag string

// Enum values for StatusFlag
const (
	StatusFlagPendingAcceptance          StatusFlag = "PENDING_ACCEPTANCE"
	StatusFlagPendingCustomerAction      StatusFlag = "PENDING_CUSTOMER_ACTION"
	StatusFlagPendingAuthorization       StatusFlag = "PENDING_AUTHORIZATION"
	StatusFlagPendingPaymentVerification StatusFlag = "PENDING_PAYMENT_VERIFICATION"
	StatusFlagPendingSupportCase         StatusFlag = "PENDING_SUPPORT_CASE"
)

// Values returns all known values for StatusFlag. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StatusFlag) Values() []StatusFlag {
	return []StatusFlag{
		"PENDING_ACCEPTANCE",
		"PENDING_CUSTOMER_ACTION",
		"PENDING_AUTHORIZATION",
		"PENDING_PAYMENT_VERIFICATION",
		"PENDING_SUPPORT_CASE",
	}
}

type Transferable string

// Enum values for Transferable
const (
	TransferableTransferable           Transferable = "TRANSFERABLE"
	TransferableUntransferable         Transferable = "UNTRANSFERABLE"
	TransferableDontKnow               Transferable = "DONT_KNOW"
	TransferableDomainInOwnAccount     Transferable = "DOMAIN_IN_OWN_ACCOUNT"
	TransferableDomainInAnotherAccount Transferable = "DOMAIN_IN_ANOTHER_ACCOUNT"
	TransferablePremiumDomain          Transferable = "PREMIUM_DOMAIN"
)

// Values returns all known values for Transferable. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Transferable) Values() []Transferable {
	return []Transferable{
		"TRANSFERABLE",
		"UNTRANSFERABLE",
		"DONT_KNOW",
		"DOMAIN_IN_OWN_ACCOUNT",
		"DOMAIN_IN_ANOTHER_ACCOUNT",
		"PREMIUM_DOMAIN",
	}
}
