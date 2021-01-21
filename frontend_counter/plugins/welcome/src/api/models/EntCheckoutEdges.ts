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
    EntCheckIn,
    EntCheckInFromJSON,
    EntCheckInFromJSONTyped,
    EntCheckInToJSON,
    EntCounterStaff,
    EntCounterStaffFromJSON,
    EntCounterStaffFromJSONTyped,
    EntCounterStaffToJSON,
    EntStatus,
    EntStatusFromJSON,
    EntStatusFromJSONTyped,
    EntStatusToJSON,
    EntStatusOpinion,
    EntStatusOpinionFromJSON,
    EntStatusOpinionFromJSONTyped,
    EntStatusOpinionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCheckoutEdges
 */
export interface EntCheckoutEdges {
    /**
     * 
     * @type {EntCheckIn}
     * @memberof EntCheckoutEdges
     */
    checkins?: EntCheckIn;
    /**
     * 
     * @type {EntCounterStaff}
     * @memberof EntCheckoutEdges
     */
    counterstaffs?: EntCounterStaff;
    /**
     * 
     * @type {EntStatusOpinion}
     * @memberof EntCheckoutEdges
     */
    statusopinion?: EntStatusOpinion;
    /**
     * 
     * @type {EntStatus}
     * @memberof EntCheckoutEdges
     */
    statuss?: EntStatus;
}

export function EntCheckoutEdgesFromJSON(json: any): EntCheckoutEdges {
    return EntCheckoutEdgesFromJSONTyped(json, false);
}

export function EntCheckoutEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCheckoutEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkins': !exists(json, 'checkins') ? undefined : EntCheckInFromJSON(json['checkins']),
        'counterstaffs': !exists(json, 'counterstaffs') ? undefined : EntCounterStaffFromJSON(json['counterstaffs']),
        'statusopinion': !exists(json, 'statusopinion') ? undefined : EntStatusOpinionFromJSON(json['statusopinion']),
        'statuss': !exists(json, 'statuss') ? undefined : EntStatusFromJSON(json['statuss']),
    };
}

export function EntCheckoutEdgesToJSON(value?: EntCheckoutEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkins': EntCheckInToJSON(value.checkins),
        'counterstaffs': EntCounterStaffToJSON(value.counterstaffs),
        'statusopinion': EntStatusOpinionToJSON(value.statusopinion),
        'statuss': EntStatusToJSON(value.statuss),
    };
}


