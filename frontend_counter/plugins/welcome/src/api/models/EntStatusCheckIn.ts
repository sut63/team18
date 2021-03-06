/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntStatusCheckInEdges,
    EntStatusCheckInEdgesFromJSON,
    EntStatusCheckInEdgesFromJSONTyped,
    EntStatusCheckInEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusCheckIn
 */
export interface EntStatusCheckIn {
    /**
     * 
     * @type {EntStatusCheckInEdges}
     * @memberof EntStatusCheckIn
     */
    edges?: EntStatusCheckInEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStatusCheckIn
     */
    id?: number;
    /**
     * StatusName holds the value of the "status_name" field.
     * @type {string}
     * @memberof EntStatusCheckIn
     */
    statusName?: string;
}

export function EntStatusCheckInFromJSON(json: any): EntStatusCheckIn {
    return EntStatusCheckInFromJSONTyped(json, false);
}

export function EntStatusCheckInFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusCheckIn {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntStatusCheckInEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'statusName': !exists(json, 'status_name') ? undefined : json['status_name'],
    };
}

export function EntStatusCheckInToJSON(value?: EntStatusCheckIn | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntStatusCheckInEdgesToJSON(value.edges),
        'id': value.id,
        'status_name': value.statusName,
    };
}


