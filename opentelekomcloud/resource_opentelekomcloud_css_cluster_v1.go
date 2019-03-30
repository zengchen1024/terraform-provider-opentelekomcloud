// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package opentelekomcloud

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceCssClusterV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceCssClusterV1Create,
		Read:   resourceCssClusterV1Read,
		Update: resourceCssClusterV1Update,
		Delete: resourceCssClusterV1Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"initial_node_num": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"node_config": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flavor": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"network_info": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"vpc_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"volume": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"encryption_key": {
										Type:     schema.TypeString,
										Computed: true,
										Optional: true,
										ForceNew: true,
									},
									"size": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"volume_type": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"availability_zone": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"add_node_num": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"enable_https": {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"datastore": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCssClusterV1UserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"add_node_num":     d.Get("add_node_num"),
		"enable_https":     d.Get("enable_https"),
		"initial_node_num": d.Get("initial_node_num"),
		"name":             d.Get("name"),
		"node_config":      d.Get("node_config"),
	}
}

func resourceCssClusterV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceCssClusterV1UserInputParams(d)

	arrayIndex := map[string]int{
		"node_config.network_info": 0,
		"node_config.volume":       0,
		"node_config":              0,
	}

	r, err := sendCssClusterV1CreateRequest(d, opts, arrayIndex, client)
	if err != nil {
		return fmt.Errorf("Error creating CssClusterV1: %s", err)
	}

	obj, err := asyncWaitCssClusterV1Create(d, config, r, client)
	if err != nil {
		return err
	}
	id, err := navigateValue(obj, []string{"cluster", "id"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id.(string))

	return resourceCssClusterV1Read(d, meta)
}

func resourceCssClusterV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	res := make(map[string]interface{})

	err = readCssClusterV1Read(d, client, res)
	if err != nil {
		return err
	}

	return setCssClusterV1Properties(d, res)
}

func resourceCssClusterV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceCssClusterV1UserInputParams(d)

	arrayIndex := map[string]int{
		"node_config.network_info": 0,
		"node_config.volume":       0,
		"node_config":              0,
	}

	_, err = sendCssClusterV1ExtendClusterRequest(d, opts, arrayIndex, client)
	if err != nil {
		return err
	}

	return resourceCssClusterV1Read(d, meta)
}

func resourceCssClusterV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "css", serviceProjectLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONBody:     nil,
		JSONResponse: &r.Body,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Cluster %q: %s", d.Id(), r.Err)
	}

	_, err = waitToFinish(
		[]string{"Done"}, []string{"Pending"},
		d.Timeout(schema.TimeoutCreate),
		1*time.Second,
		func() (interface{}, string, error) {
			resp, err := client.Get(
				url, nil,
				&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return resp, "Done", nil
				}
				return nil, "", nil
			}
			return resp, "Pending", nil
		},
	)
	return err
}

func sendCssClusterV1CreateRequest(d *schema.ResourceData, opts map[string]interface{},
	arrayIndex map[string]int, client *golangsdk.ServiceClient) (interface{}, error) {
	url := client.ServiceURL("clusters")

	params, err := buildCssClusterV1CreateParameters(opts, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error building the request body of api(create)")
	}

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error run api(create): %s", r.Err)
	}
	return r.Body, nil
}

func buildCssClusterV1CreateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	diskEncryptionProp, err := expandCssClusterV1CreateDiskEncryption(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(diskEncryptionProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["diskEncryption"] = diskEncryptionProp
	}

	httpsEnableProp, err := expandCssClusterV1CreateHttpsEnable(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(httpsEnableProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["httpsEnable"] = httpsEnableProp
	}

	instanceProp, err := expandCssClusterV1CreateInstance(opts, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(instanceProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["instance"] = instanceProp
	}

	instanceNumProp, err := navigateValue(opts, []string{"initial_node_num"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(instanceNumProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["instanceNum"] = instanceNumProp
	}

	nameProp, err := navigateValue(opts, []string{"name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(nameProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["name"] = nameProp
	}

	params = map[string]interface{}{"cluster": params}

	return params, nil
}

func expandCssClusterV1CreateDiskEncryption(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	systemCmkidProp, err := navigateValue(d, []string{"node_config", "volume", "encryption_key"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(systemCmkidProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["systemCmkid"] = systemCmkidProp
	}

	systemEncryptedProp, err := expandCssClusterV1CreateDiskEncryptionSystemEncrypted(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(systemEncryptedProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["systemEncrypted"] = systemEncryptedProp
	}

	return req, nil
}

func expandCssClusterV1CreateDiskEncryptionSystemEncrypted(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	v, err := navigateValue(d, []string{"node_config", "volume", "encryption_key"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if v1, ok := v.(string); ok && v1 != "" {
		return "1", nil
	}
	return "0", nil
}

func expandCssClusterV1CreateHttpsEnable(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	v, err := navigateValue(d, []string{"enable_https"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if v1, ok := v.(bool); ok && v1 {
		return "true", nil
	}
	return "false", nil
}

func expandCssClusterV1CreateInstance(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	availabilityZoneProp, err := navigateValue(d, []string{"node_config", "availability_zone"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(availabilityZoneProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["availability_zone"] = availabilityZoneProp
	}

	flavorRefProp, err := navigateValue(d, []string{"node_config", "flavor"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(flavorRefProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["flavorRef"] = flavorRefProp
	}

	nicsProp, err := expandCssClusterV1CreateInstanceNics(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(nicsProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["nics"] = nicsProp
	}

	volumeProp, err := expandCssClusterV1CreateInstanceVolume(d, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(volumeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["volume"] = volumeProp
	}

	return req, nil
}

func expandCssClusterV1CreateInstanceNics(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	netIdProp, err := navigateValue(d, []string{"node_config", "network_info", "subnet_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(netIdProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["netId"] = netIdProp
	}

	securityGroupIdProp, err := navigateValue(d, []string{"node_config", "network_info", "security_group_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(securityGroupIdProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["securityGroupId"] = securityGroupIdProp
	}

	vpcIdProp, err := navigateValue(d, []string{"node_config", "network_info", "vpc_id"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(vpcIdProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["vpcId"] = vpcIdProp
	}

	return req, nil
}

func expandCssClusterV1CreateInstanceVolume(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	req := make(map[string]interface{})

	sizeProp, err := navigateValue(d, []string{"node_config", "volume", "size"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(sizeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["size"] = sizeProp
	}

	volumeTypeProp, err := navigateValue(d, []string{"node_config", "volume", "volume_type"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err = isEmptyValue(reflect.ValueOf(volumeTypeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		req["volume_type"] = volumeTypeProp
	}

	return req, nil
}

func asyncWaitCssClusterV1Create(d *schema.ResourceData,
	config *Config, result interface{}, client *golangsdk.ServiceClient) (interface{}, error) {

	pathParameters := map[string][]string{
		"id": []string{"cluster", "id"},
	}
	var data = make(map[string]string)
	for key, path := range pathParameters {
		value, err := navigateValue(result, path, nil)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving async operation path parameter: %s", err)
		}
		data[key] = value.(string)
	}
	url, err := replaceVars(d, "clusters/{id}", data)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"200"},
		[]string{"100"},
		d.Timeout(schema.TimeoutCreate),
		1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(
				url, &r.Body,
				&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			status, err := navigateValue(r.Body, []string{"cluster", "status"}, nil)
			if err != nil {
				return nil, "", nil
			}
			return r.Body, status.(string), nil
		},
	)
}

func sendCssClusterV1ExtendClusterRequest(d *schema.ResourceData, opts map[string]interface{},
	arrayIndex map[string]int, client *golangsdk.ServiceClient) (interface{}, error) {
	url, err := replaceVars(d, "clusters/{id}/extend", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	params, err := buildCssClusterV1ExtendClusterParameters(opts, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error building the request body of api(extend_cluster)")
	}

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error run api(extend_cluster): %s", r.Err)
	}
	return r.Body, nil
}

func buildCssClusterV1ExtendClusterParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	modifySizeProp, err := navigateValue(opts, []string{"add_node_num"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	e, err := isEmptyValue(reflect.ValueOf(modifySizeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		params["modifySize"] = modifySizeProp
	}

	params = map[string]interface{}{"grow": params}

	return params, nil
}

func asyncWaitCssClusterV1ExtendCluster(d *schema.ResourceData,
	config *Config, result interface{}, client *golangsdk.ServiceClient) (interface{}, error) {

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return nil, err
	}
	url = client.ServiceURL(url)

	return waitToFinish(
		[]string{"200"},
		[]string{"100"},
		d.Timeout(schema.TimeoutUpdate),
		1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(
				url, &r.Body,
				&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			status, err := navigateValue(r.Body, []string{"cluster", "status"}, nil)
			if err != nil {
				return nil, "", nil
			}
			return r.Body, status.(string), nil
		},
	)
}

func readCssClusterV1Read(d *schema.ResourceData, client *golangsdk.ServiceClient, result map[string]interface{}) error {
	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(
		url, &r.Body,
		&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return fmt.Errorf("Error running api(read) for resource(CssClusterV1: %v), error: %s", d.Id(), r.Err)
	}

	result["read"] = r.Body

	return nil
}

func setCssClusterV1Properties(d *schema.ResourceData, response map[string]interface{}) error {
	opts := resourceCssClusterV1UserInputParams(d)

	actionsProp, err := navigateValue(response, []string{"read", "actions"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:actions, err: %s", err)
	}
	if err = d.Set("actions", actionsProp); err != nil {
		return fmt.Errorf("Error setting Cluster:actions, err: %s", err)
	}

	createdProp, err := navigateValue(response, []string{"read", "created"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:created, err: %s", err)
	}
	if err = d.Set("created", createdProp); err != nil {
		return fmt.Errorf("Error setting Cluster:created, err: %s", err)
	}

	datastoreProp, _ := opts["datastore"]
	datastoreProp, err = flattenCssClusterV1Datastore(response, nil, datastoreProp)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:datastore, err: %s", err)
	}
	if err = d.Set("datastore", datastoreProp); err != nil {
		return fmt.Errorf("Error setting Cluster:datastore, err: %s", err)
	}

	enableHttpsProp, err := navigateValue(response, []string{"read", "httpsEnable"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:enable_https, err: %s", err)
	}
	if err = d.Set("enable_https", enableHttpsProp); err != nil {
		return fmt.Errorf("Error setting Cluster:enable_https, err: %s", err)
	}

	endpointProp, err := navigateValue(response, []string{"read", "endpoint"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:endpoint, err: %s", err)
	}
	if err = d.Set("endpoint", endpointProp); err != nil {
		return fmt.Errorf("Error setting Cluster:endpoint, err: %s", err)
	}

	nameProp, err := navigateValue(response, []string{"read", "name"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:name, err: %s", err)
	}
	if err = d.Set("name", nameProp); err != nil {
		return fmt.Errorf("Error setting Cluster:name, err: %s", err)
	}

	nodeConfigProp, _ := opts["node_config"]
	nodeConfigProp, err = flattenCssClusterV1NodeConfig(response, nil, nodeConfigProp)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:node_config, err: %s", err)
	}
	if err = d.Set("node_config", nodeConfigProp); err != nil {
		return fmt.Errorf("Error setting Cluster:node_config, err: %s", err)
	}

	nodesProp, _ := opts["nodes"]
	nodesProp, err = flattenCssClusterV1Nodes(response, nil, nodesProp)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:nodes, err: %s", err)
	}
	if err = d.Set("nodes", nodesProp); err != nil {
		return fmt.Errorf("Error setting Cluster:nodes, err: %s", err)
	}

	updatedProp, err := navigateValue(response, []string{"read", "updated"}, nil)
	if err != nil {
		return fmt.Errorf("Error reading Cluster:updated, err: %s", err)
	}
	if err = d.Set("updated", updatedProp); err != nil {
		return fmt.Errorf("Error setting Cluster:updated, err: %s", err)
	}

	return nil
}

func flattenCssClusterV1Datastore(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	if result[0] == nil {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	typeProp, err := navigateValue(d, []string{"read", "datastore", "type"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:type, err: %s", err)
	}
	r["type"] = typeProp

	versionProp, err := navigateValue(d, []string{"read", "datastore", "version"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:version, err: %s", err)
	}
	r["version"] = versionProp

	return result, nil
}

func flattenCssClusterV1NodeConfig(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	if result[0] == nil {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	networkInfoProp, _ := r["network_info"]
	networkInfoProp, err := flattenCssClusterV1NodeConfigNetworkInfo(d, arrayIndex, networkInfoProp)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:network_info, err: %s", err)
	}
	r["network_info"] = networkInfoProp

	volumeProp, _ := r["volume"]
	volumeProp, err = flattenCssClusterV1NodeConfigVolume(d, arrayIndex, volumeProp)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:volume, err: %s", err)
	}
	r["volume"] = volumeProp

	return result, nil
}

func flattenCssClusterV1NodeConfigNetworkInfo(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	if result[0] == nil {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	securityGroupIdProp, err := navigateValue(d, []string{"read", "securityGroupId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:security_group_id, err: %s", err)
	}
	r["security_group_id"] = securityGroupIdProp

	subnetIDProp, err := navigateValue(d, []string{"read", "subnetId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:subnet_id, err: %s", err)
	}
	r["subnet_id"] = subnetIDProp

	vpcIdProp, err := navigateValue(d, []string{"read", "vpcId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:vpc_id, err: %s", err)
	}
	r["vpc_id"] = vpcIdProp

	return result, nil
}

func flattenCssClusterV1NodeConfigVolume(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		result = make([]interface{}, 1, 1)
	}
	if result[0] == nil {
		result[0] = make(map[string]interface{})
	}
	r := result[0].(map[string]interface{})

	encryptionKeyProp, err := navigateValue(d, []string{"read", "cmkId"}, arrayIndex)
	if err != nil {
		return nil, fmt.Errorf("Error reading Cluster:encryption_key, err: %s", err)
	}
	r["encryption_key"] = encryptionKeyProp

	return result, nil
}

func flattenCssClusterV1Nodes(d interface{}, arrayIndex map[string]int, currentValue interface{}) (interface{}, error) {
	result, ok := currentValue.([]interface{})
	if !ok || len(result) == 0 {
		v, err := navigateValue(d, []string{"read", "instances"}, arrayIndex)
		if err != nil {
			return nil, err
		}
		n := len(v.([]interface{}))
		result = make([]interface{}, n, n)
	}

	newArrayIndex := make(map[string]int)
	if arrayIndex != nil {
		for k, v := range arrayIndex {
			newArrayIndex[k] = v
		}
	}

	for i := 0; i < len(result); i++ {
		newArrayIndex["read.instances"] = i
		if result[i] == nil {
			result[i] = make(map[string]interface{})
		}
		r := result[i].(map[string]interface{})

		idProp, err := navigateValue(d, []string{"read", "instances", "id"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:id, err: %s", err)
		}
		r["id"] = idProp

		nameProp, err := navigateValue(d, []string{"read", "instances", "name"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:name, err: %s", err)
		}
		r["name"] = nameProp

		typeProp, err := navigateValue(d, []string{"read", "instances", "type"}, newArrayIndex)
		if err != nil {
			return nil, fmt.Errorf("Error reading Cluster:type, err: %s", err)
		}
		r["type"] = typeProp
	}

	return result, nil
}
