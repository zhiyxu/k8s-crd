// +<tag_name>[=value]
// 为整个v1包里的所有类型定义自动生成DeepCopy方法
// +k8s:deepcopy-gen=package

// 定义了这个包对应的API组的名字
// +groupName=samplecrd.k8s.io
package v1

// 定义在doc.go文件中，起到的是全局的代码生成控制的作用，所以也被称为Global Tags
