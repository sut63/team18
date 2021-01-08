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
    EntCheckInEdges,
    EntCheckInEdgesFromJSON,
    EntCheckInEdgesFromJSONTyped,
    EntCheckInEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCheckIn
 */
export interface EntCheckIn {
    /**
     * CheckinDate holds the value of the "checkin_date" field.
     * @type {string}
     * @memberof EntCheckIn
     */
    checkinDate?: string;
    /**
     * 
     * @type {EntCheckInEdges}
     * @memberof EntCheckIn
     */
    edges?: EntCheckInEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCheckIn
     */
    id?: number;
}

export function EntCheckInFromJSON(json: any): EntCheckIn {
    return EntCheckInFromJSONTyped(json, false);
}

export function EntCheckInFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCheckIn {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkinDate': !exists(json, 'checkin_date') ? undefined : json['checkin_date'],
        'edges': !exists(json, 'edges') ? undefined : EntCheckInEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCheckInToJSON(value?: EntCheckIn | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkin_date': value.checkinDate,
        'edges': EntCheckInEdgesToJSON(value.edges),
        'id': value.id,
    };
}

