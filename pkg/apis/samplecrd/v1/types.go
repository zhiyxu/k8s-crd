// 定义一个Network类型到底有哪些字段，比如spec字段里的内容
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// genclient的意思是请为下面这个API资源类型生成对应的Client代码
// 而noStatus的意思是，这个API资源类型的定义里，没有Status字段，
// 否则，生成的Client就会自动带上UpdateStatus方法
// genclient只需要写在Network类型上，而不用写在NetworkList上，
// 因为NetworkList只是一个返回值类型，Network才是主类型

//由于在Global Tags里已经定义了为所有类型生成DeepCopy方法，所以这里就不用显式地加上k8s:deepcopy-gen=true

// +genclient
// +genclient:noStatus
// 在生成DeepCopy的时候，实现Kubernetes提供的runtime.Object接口
// 否则在某些版本的Kubernetes里，这个类型定义会出现编译错误，可以当作固定套路
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network describes a Network resource
type Network struct {
	// 标准的Kubernetes对象定义：TypeMeta和ObjectMeta
	// TypeMeta is the metadata for the resource, like kind and apiversion
	// API元数据
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	// 对象元数据
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec NetworkSpec `json:"spec"`
}

// NetworkSpec is the spec for a Network resource
type NetworkSpec struct {
	// Cidr and Gateway are example custom spec fields
	//
	// this is where you would put your custom resource data
	// 这里json:后面的名字指的是这个字段被转换成JSON格式之后的名字，也就是YAML文件里的字段名字
	Cidr    string `json:"cidr"`
	Gateway string `json:"gateway"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a list of Network resources
// 用于描述一组Network对象应该包括哪些字段
// 之所以需要这样一个类型，是因为在Kubernetes中，获取所有X对象的List()方法，返回值都是List类型，而不是X类型的数组
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Network `json:"items"`
}
