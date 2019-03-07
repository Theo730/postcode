<?php
/**
  Created by Theo730
  Date: 07.03.2019 
*/

namespace yamalValidator;

use \Bitrix\Main\EventManager;

class FormValidatorPostCode{
    function getDescription(){
        return array(
            'NAME' => 'PostCode',
            'DESCRIPTION' => 'почтовый индекс',
            'TYPES' => array('text'),
            'SETTINGS' => array(__CLASS__, 'getSettings'),
            'CONVERT_TO_DB' => array(__CLASS__, 'toDB'), 
            'CONVERT_FROM_DB' => array(__CLASS__, 'fromDB'), 
            'HANDLER' => array(__CLASS__, 'doValidate')
        );
    }
    function getSettings(){
        return array();
    }
    function toDB($arParams){
        return serialize($arParams);
    }
    function fromDB($strParams){
        return unserialize($strParams);
    }
    function getRest(int $postIdex){
    	    $serverURI		= \COption::GetOptionString("<module>", 'POSTCODE_URI_API', '');
	    $serverLogin	= \COption::GetOptionString("<modeule>", 'POSTCODE_LOGIN', '');
    	    $serverPassword	= \COption::GetOptionString("<module>", 'POSTCODE_PASSWORD', '');
/// лучше поместить в библиотеку
	    $base64		= base64_encode($serverLogin.':'.$serverPassword);
	    $opts	= array( 
		'http'	=> array(
		    'header'		=> 'Authorization: Basic ' . $base64,
		    'method'		=> 'GET',
		),
		'ssl'	=> array( // tckb https , то без проверок
		    'verify_peer'	=> false,
		    'verify_peer_name'	=> false,
		),
	    );
	    $context = stream_context_create($opts);
	    $result = file_get_contents($serverURI."validatePostIndex/".$postIndex , false, $context);
	return intval($result);
    }
    function doValidate($arParams, $arQuestion, $arAnswers, $arValues){
        global $APPLICATION;
	if((\COption::GetOptionString("<module>", 'POSTCODE_ENABLE', 'Y') == 'Y')&&(\Bitrix\Main\Loader::includeModule("<module>"))){
            foreach ($arValues as $value){
		if(strlen(trim($value))==0) continue;
        	if(getRest(intval($value))<1){
		    if(\COption::GetOptionString("<module>", 'SYSLOG_CHECK', 'Y') == 'Y')
			toSyslog($arQuestion['TITLE'].'('.$value.')');
        	    }
		    $APPLICATION->ThrowException('Веденного почтового индекса не существует '.$arQuestion['TITLE'].'');
        	    return false;
    	    }
    	}
	return true;
    }
}
