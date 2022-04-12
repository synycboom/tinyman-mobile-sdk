package com.example.tinyman

import android.os.Bundle
import android.util.Log
import com.google.android.material.snackbar.Snackbar
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.tinyman.databinding.FragmentRedeemBinding

class RedeemFragment : Fragment() {

    private var _binding: FragmentRedeemBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        _binding = FragmentRedeemBinding.inflate(inflater, container, false)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        binding.textviewOutput.text = "This example redeem excess tokens from swapping"
        binding.buttonRedeem.setOnClickListener {
            val asset1Id = binding.editTextAsset1Id.text
            if (asset1Id.isEmpty()) {
                Snackbar.make(view, "Asset 1 ID is empty", Snackbar.LENGTH_LONG)
                    .setAction("Action", null).show()
            } else {
                binding.textviewOutput.text = "Loading..."
                binding.buttonRedeem.isEnabled = false
                binding.buttonRedeem.isClickable = false
                try {
                    SDKViewModel().redeem(asset1Id.toString()) {
                        binding.buttonRedeem.isEnabled = true
                        binding.buttonRedeem.isClickable = true
                        Log.i("TINY_MAN_MOBILE_SDK", it)
                        binding.textviewOutput.text = it
                    }
                } catch (err: Throwable) {
                    binding.textviewOutput.text = err.stackTraceToString()
                }
            }
        }
        binding.buttonGoBack.setOnClickListener {
            findNavController().navigate(R.id.action_RedeemFragment_to_ExampleFragment)
        }
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}