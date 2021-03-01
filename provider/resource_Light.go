package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLight() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLightCreate,
		ReadContext:   resourceLightRead,
		UpdateContext: resourceLightUpdate,
		DeleteContext: resourceLightDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "on",
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Light",
			},
			"color_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "hs",
			},
			"brightness": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  255,
			},
			"white_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"supported_features": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  147,
			},
			"hs_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"rgb_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"xy_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	}
}
func resourceLightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics
	c := m.(*Client)
	//return diag.FromErr(fmt.Errorf("%[1]T: %[1]v", d.Get("rgb_color.#")))
	item := LightItem{
		EntityID: d.Get("entity_id").(string),
		State:    d.Get("state").(string),
		Attr: Attributes{
			Brightness:        d.Get("brightness").(int),
			WhiteValue:        d.Get("white_value").(int),
			Name:              d.Get("friendly_name").(string),
			ColorMode:         d.Get("color_mode").(string),
			SupportedFeatures: d.Get("supported_features").(int),
		},
	}
	if d.Get("hs_color.#").(int) != 0 {
		item.Attr.HsColor = []float64{d.Get("hs_color.0").(float64), d.Get("hs_color.1").(float64)}
	}
	if d.Get("rgb_color.#").(int) != 0 {
		item.Attr.RgbColor = []int{d.Get("rgb_color.0").(int), d.Get("rgb_color.1").(int), d.Get("rgb_color.2").(int)}
	}
	if d.Get("xy_color.#").(int) != 0 {
		item.Attr.XyColor = []float64{d.Get("xy_color.0").(float64), d.Get("xy_color.1").(float64)}
	}
	//			HsColor:           d.Get("hs_color").([]float64),
	//RgbColor:          d.Get("rgb_color").([]int),
	//XyColor:           d.Get("xy_color").([]float64),
	o, err := c.StartLight(item)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(o.EntityID)
	return resourceLightRead(ctx, d, m)
	//return diags
}
func resourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	lightID := d.Id()

	light, err := c.GetLight(lightID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", light.State); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hs_color", light.Attr.HsColor); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rgb_color", light.Attr.RgbColor); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("xy_color", light.Attr.XyColor); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
func resourceLightUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	item := LightItem{
		EntityID: d.Get("entity_id").(string),
		State:    d.Get("state").(string),
		Attr: Attributes{
			Brightness:        d.Get("brightness").(int),
			WhiteValue:        d.Get("white_value").(int),
			Name:              d.Get("friendly_name").(string),
			ColorMode:         d.Get("color_mode").(string),
			SupportedFeatures: d.Get("supported_features").(int),
		},
	}
	if d.Get("hs_color.#").(int) != 0 {
		item.Attr.HsColor = []float64{d.Get("hs_color.0").(float64), d.Get("hs_color.1").(float64)}
	}
	if d.Get("rgb_color.#").(int) != 0 {
		item.Attr.RgbColor = []int{d.Get("rgb_color.0").(int), d.Get("rgb_color.1").(int), d.Get("rgb_color.2").(int)}
	}
	if d.Get("xy_color.#").(int) != 0 {
		item.Attr.XyColor = []float64{d.Get("xy_color.0").(float64), d.Get("xy_color.1").(float64)}
	}
	_, err := c.StartLight(item)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceLightRead(ctx, d, m)
}
func resourceLightDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	lightID := d.Id()

	err := c.DelLight(lightID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
